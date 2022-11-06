package user

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/ahmadlubis/lavandeapp/module/backend/utility"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"regexp"
)

const asciiRegex = "^[\\x00-\\x7F]+$"

type userRegistrationUsecase struct {
	db *gorm.DB
}

func NewUserRegistrationUsecase(db *gorm.DB) usecase.UserRegistrationUsecase {
	return &userRegistrationUsecase{db: db}
}

func (u *userRegistrationUsecase) Register(_ context.Context, req request.RegisterUserRequest) (entity.User, error) {
	if err := ValidateUserPassword(req.Password, req.Email); err != nil {
		return entity.User{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, model.NewUnknownError(req.Email, err)
	}

	user := entity.User{
		Name:     req.Name,
		NIK:      req.NIK,
		Email:    req.Email,
		PhoneNo:  req.PhoneNo,
		Role:     entity.UserRoleResident,
		Status:   entity.UserStatusActive,
		Password: passwordHash,
	}
	if err = user.Validate(); err != nil {
		return entity.User{}, err
	}

	if result := u.db.Create(&user); result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == utility.MysqlErrorConflictNumber {
				return entity.User{}, model.NewExpectedError("another account with the same NIK / Email / PhoneNo already exists", "USER_CONFLICT", http.StatusConflict, req.Email)
			}
		}
		return entity.User{}, model.NewUnknownError(req.Email, result.Error)
	}

	return user, nil
}

func ValidateUserPassword(passwd, trackId string) error {
	if len(passwd) < 8 || len(passwd) > 32 {
		return model.NewExpectedError("password must be between 8 to 32 characters long", "USER_INVALID", http.StatusBadRequest, trackId)
	}
	if match, _ := regexp.MatchString(asciiRegex, passwd); !match {
		return model.NewExpectedError("password can't contains non-standard characters", "USER_INVALID", http.StatusBadRequest, trackId)
	}
	return nil
}

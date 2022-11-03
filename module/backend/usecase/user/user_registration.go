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

const (
	numbersOnlyRegex = "^[0-9]*$"
	emailRegex       = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	phoneNoRegex     = "^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$"
	asciiRegex       = "^[\\x00-\\x7F]+$"
)

type userRegistrationUsecase struct {
	db *gorm.DB
}

func NewUserRegistrationUsecase(db *gorm.DB) usecase.UserRegistrationUsecase {
	return &userRegistrationUsecase{db: db}
}

func (u *userRegistrationUsecase) RegisterUser(_ context.Context, req request.RegisterUserRequest) (entity.User, error) {
	var err = u.validateRegisterRequest(req)
	if err != nil {
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
	if result := u.db.Create(&user); result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == utility.MysqlErrorConflictNumber {
				return entity.User{}, model.NewExpectedError("another account with the same NIK / Email / PhoneNo already exists", "USER_CONFLICT", http.StatusConflict, req.Email)
			}
		}
		return entity.User{}, model.NewUnknownError(req.Email, err)
	}

	return user, nil
}

func (u *userRegistrationUsecase) validateRegisterRequest(req request.RegisterUserRequest) error {
	if len(req.Email) == 0 || len(req.Email) > 255 {
		return model.NewExpectedError("email must be present and be at most 255 characters long", "USER_INVALID", http.StatusBadRequest, req.Email)
	}
	if match, _ := regexp.MatchString(emailRegex, req.Email); !match {
		return model.NewExpectedError("email must be a valid email address", "USER_INVALID", http.StatusBadRequest, req.Email)
	}
	if len(req.Name) > 255 {
		return model.NewExpectedError("name must be at most 255 characters long", "USER_INVALID", http.StatusBadRequest, req.Email)
	}
	if len(req.NIK) > 0 {
		if len(req.NIK) != 16 {
			return model.NewExpectedError("NIK must be 16 characters long", "USER_INVALID", http.StatusBadRequest, req.Email)
		}
		if match, _ := regexp.MatchString(numbersOnlyRegex, req.NIK); !match {
			return model.NewExpectedError("NIK must only consist of numbers", "USER_INVALID", http.StatusBadRequest, req.Email)
		}
	}
	if len(req.PhoneNo) > 0 {
		if match, _ := regexp.MatchString(phoneNoRegex, req.PhoneNo); !match {
			return model.NewExpectedError("phone_no must be a valid phone number", "USER_INVALID", http.StatusBadRequest, req.Email)
		}
	}
	if len(req.Password) < 8 || len(req.Password) > 32 {
		return model.NewExpectedError("password must be between 8 to 32 characters long", "USER_INVALID", http.StatusBadRequest, req.Email)
	}
	if match, _ := regexp.MatchString(asciiRegex, req.Password); !match {
		return model.NewExpectedError("password can't contains non-standard characters", "USER_INVALID", http.StatusBadRequest, req.Email)
	}

	return nil
}

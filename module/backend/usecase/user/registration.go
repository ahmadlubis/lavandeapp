package user

import (
	"context"
	"net/http"

	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/ahmadlubis/lavandeapp/module/backend/utility"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRegistrationUsecase struct {
	db *gorm.DB
}

func NewUserRegistrationUsecase(db *gorm.DB) usecase.UserRegistrationUsecase {
	return &userRegistrationUsecase{db: db}
}

func (u *userRegistrationUsecase) Register(_ context.Context, req request.RegisterUserRequest) (entity.User, error) {
	if err := utility.ValidateUserPassword(req.Password, req.Email); err != nil {
		return entity.User{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return entity.User{}, model.NewUnknownError(req.Email, err)
	}

	religion, err := entity.ParseUserReligion(req.Religion)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Name:     req.Name,
		NIK:      req.NIK,
		Email:    req.Email,
		PhoneNo:  req.PhoneNo,
		Role:     entity.UserRoleResident,
		Status:   entity.UserStatusActive,
		Religion: religion,
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

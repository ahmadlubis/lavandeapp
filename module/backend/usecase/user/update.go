package user

import (
	"context"
	"errors"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/ahmadlubis/lavandeapp/module/backend/utility"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type userUpdateUsecase struct {
	db *gorm.DB
}

func NewUserUpdateUsecase(db *gorm.DB) usecase.UserUpdateUsecase {
	return &userUpdateUsecase{db: db}
}

func (u *userUpdateUsecase) Update(_ context.Context, req request.UpdateUserRequest) (entity.User, error) {
	var err error
	var user entity.User
	if err = u.db.Where("email = ?", req.TargetEmail).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, model.InvalidTokenError.WithTrackId(req.TargetEmail)
		}
		return entity.User{}, model.NewUnknownError(req.TargetEmail, err)
	}

	if req.Password != "" {
		if err := ValidateUserPassword(req.Password, req.Email); err != nil {
			return entity.User{}, err
		}

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return entity.User{}, model.NewUnknownError(req.Email, err)
		}

		user.Password = passwordHash
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.NIK != "" {
		user.NIK = req.NIK
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.PhoneNo != "" {
		user.PhoneNo = req.PhoneNo
	}
	if req.Religion != "" {
		user.Religion, err = entity.ParseUserReligion(req.Religion)
		if err != nil {
			return entity.User{}, err
		}
	}

	if err := user.Validate(); err != nil {
		return entity.User{}, err
	}
	if result := u.db.Save(&user); result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == utility.MysqlErrorConflictNumber {
				return entity.User{}, model.NewExpectedError("another account with the same NIK / Email / PhoneNo already exists", "USER_CONFLICT", http.StatusConflict, req.Email)
			}
		}
		return entity.User{}, model.NewUnknownError(req.Email, result.Error)
	}

	return user, nil
}
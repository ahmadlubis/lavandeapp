package user

import (
	"context"
	"errors"
	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type userLoginUsecase struct {
	cfg config.AuthConfig
	db  *gorm.DB
}

func NewUserLoginUsecase(cfg config.AuthConfig, db *gorm.DB) usecase.UserLoginUsecase {
	return &userLoginUsecase{cfg: cfg, db: db}
}

func (u *userLoginUsecase) Login(ctx context.Context, req request.LoginUserRequest) (model.AccessToken, error) {
	var user entity.User
	var invalidLoginError = model.NewExpectedError("wrong email or password", "USER_UNAUTHORIZED", http.StatusUnauthorized, req.Email)

	if err := u.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.AccessToken{}, invalidLoginError
		}
		return model.AccessToken{}, model.NewUnknownError(req.Email, err)
	}

	if user.Status == entity.UserStatusNonactive {
		return model.AccessToken{}, model.DeactivatedAccountError.WithTrackId(req.Email)
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return model.AccessToken{}, invalidLoginError
	}

	return u.generateJwt(ctx, user)
}

func (u *userLoginUsecase) generateJwt(_ context.Context, user entity.User) (model.AccessToken, error) {
	now := time.Now()
	expiredAt := now.Add(1 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role.String(),
		"exp":  expiredAt.Unix(),
		"iat":  now.Unix(),
		"nbf":  now.Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.cfg.JWTSecretKey))
	if err != nil {
		return model.AccessToken{}, model.NewUnknownError(user.Email, err)
	}

	return model.AccessToken{
		AccessToken: tokenString,
		ExpiredAt:   expiredAt,
	}, nil
}

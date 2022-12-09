package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type userTokenVerificationUsecase struct {
	cfg config.AuthConfig
	db  *gorm.DB
}

func NewUserTokenVerificationUsecase(cfg config.AuthConfig, db *gorm.DB) usecase.UserTokenVerificationUsecase {
	return &userTokenVerificationUsecase{cfg: cfg, db: db}
}

func (u *userTokenVerificationUsecase) VerifyToken(ctx context.Context, tokenString string) (entity.User, error) {
	var id, err = u.parseAndVerifyToken(ctx, tokenString)
	if err != nil {
		return entity.User{}, model.InvalidTokenError
	}

	var user entity.User
	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, model.InvalidTokenError
		}
		return entity.User{}, model.NewUnknownError(strconv.FormatUint(id, 10), err)
	}

	if user.Status == entity.UserStatusNonactive {
		return entity.User{}, model.DeactivatedAccountError
	}
	var tenant entity.Tenant
	if err := u.db.Where("user_id = ? AND role = ?", id, entity.TenantRoleOwner).First(&tenant).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, model.NewUnknownError(strconv.FormatUint(id, 10), err)
		}
	} else {
		user.IsOwner = true
	}

	return user, nil
}

func (u *userTokenVerificationUsecase) parseAndVerifyToken(_ context.Context, tokenString string) (uint64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.InvalidTokenError
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(u.cfg.JWTSecretKey), nil
	})
	if err != nil {
		return 0, model.InvalidTokenError
	}

	if token.Claims.Valid() != nil {
		return 0, model.InvalidTokenError
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, model.InvalidTokenError
	}
	return uint64(claims["id"].(float64)), nil
}

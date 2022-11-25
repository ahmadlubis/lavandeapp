package unit

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type unitUpdateUsecase struct {
	db *gorm.DB
}

func NewUnitUpdateUsecase(db *gorm.DB) usecase.UnitUpdateUsecase {
	return &unitUpdateUsecase{db: db}
}

func (u *unitUpdateUsecase) Update(_ context.Context, req request.UpdateUnitRequest) (entity.Unit, error) {
	var unit entity.Unit
	if err := u.db.Where("id = ?", strconv.FormatUint(req.ID, 10)).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Unit{}, model.InvalidTokenError.WithTrackId(strconv.FormatUint(req.ID, 10))
		}
		return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.ID, 10), err)
	}

	var err error
	var ajb, akte []byte
	if req.GovID != nil {
		unit.GovID = *req.GovID
	}
	if req.Tower != nil {
		unit.Tower = *req.Tower
	}
	if req.Floor != nil {
		unit.Floor = *req.Floor
	}
	if req.UnitNo != nil {
		unit.UnitNo = *req.UnitNo
	}
	if req.AJB != nil {
		if ajb, err = base64.StdEncoding.DecodeString(*req.AJB); err != nil {
			return entity.Unit{}, model.NewExpectedError("ajb is not a valid base64 string", "UNIT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.ID, 10))
		}
		unit.AJB = ajb
	}
	if req.Akte != nil {
		if akte, err = base64.StdEncoding.DecodeString(*req.Akte); err != nil {
			return entity.Unit{}, model.NewExpectedError("akte is not a valid base64 string", "UNIT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.ID, 10))
		}
		unit.Akte = akte
	}

	if err := unit.Validate(); err != nil {
		return entity.Unit{}, err
	}
	if result := u.db.Save(&unit); result.Error != nil {
		return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.ID, 10), result.Error)
	}

	return unit, nil
}

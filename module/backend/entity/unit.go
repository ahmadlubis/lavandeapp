package entity

import (
	"encoding/base64"
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
	"time"
)

const fileMaxBytes = 16777215 // 16MB

type Base64EncodedFile []byte

func (f Base64EncodedFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(base64.StdEncoding.EncodeToString(f))
}

type Unit struct {
	ID        uint64            `gorm:"column:id" json:"id"`
	GovID     string            `gorm:"column:gov_id" json:"gov_id"`
	Tower     string            `gorm:"column:tower" json:"tower"`
	Floor     string            `gorm:"column:floor" json:"floor"`
	UnitNo    string            `gorm:"column:unit_no" json:"unit_no"`
	AJB       Base64EncodedFile `gorm:"column:ajb" json:"ajb"`
	Akte      Base64EncodedFile `gorm:"column:akte" json:"akte"`
	CreatedAt time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time         `gorm:"column:updated_at" json:"updated_at"`
}

func (Unit) TableName() string {
	return "units"
}

func (u Unit) Validate() error {
	if len(u.GovID) == 0 {
		return model.NewExpectedError("gov_id must be present", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	if len(u.Tower) == 0 {
		return model.NewExpectedError("tower must be present", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	if len(u.Floor) == 0 {
		return model.NewExpectedError("floor must be present", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	if len(u.UnitNo) == 0 {
		return model.NewExpectedError("unit_no must be present", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	if len(u.AJB) > fileMaxBytes {
		return model.NewExpectedError("ajb filesize can't be more than 16MB", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	if len(u.Akte) > fileMaxBytes {
		return model.NewExpectedError("akte filesize can't be more than 16MB", "UNIT_INVALID", http.StatusBadRequest, u.GovID)
	}
	return nil
}

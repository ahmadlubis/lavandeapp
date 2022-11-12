package entity

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
	"regexp"
	"time"
)

const (
	numbersOnlyRegex = "^[0-9]*$"
	emailRegex       = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	phoneNoRegex     = "^[\\+]?[(]?[0-9]{3}[)]?[-\\s\\.]?[0-9]{3}[-\\s\\.]?[0-9]{4,6}$"
)

type User struct {
	ID        uint64       `gorm:"column:id" json:"id"`
	Name      string       `gorm:"column:name" json:"name"`
	NIK       string       `gorm:"column:nik" json:"nik"`
	Email     string       `gorm:"column:email" json:"email"`
	PhoneNo   string       `gorm:"column:phone_no" json:"phone_no"`
	Role      UserRole     `gorm:"column:role" json:"role"`
	Status    UserStatus   `gorm:"column:status" json:"status"`
	Religion  UserReligion `gorm:"column:religion" json:"religion"`
	Password  []byte       `gorm:"column:password" json:"-"`
	CreatedAt time.Time    `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func (u User) Validate() error {
	if len(u.Email) == 0 || len(u.Email) > 255 {
		return model.NewExpectedError("email must be present and be at most 255 characters long", "USER_INVALID", http.StatusBadRequest, u.Email)
	}
	if match, _ := regexp.MatchString(emailRegex, u.Email); !match {
		return model.NewExpectedError("email must be a valid email address", "USER_INVALID", http.StatusBadRequest, u.Email)
	}
	if len(u.Name) > 255 {
		return model.NewExpectedError("name must be at most 255 characters long", "USER_INVALID", http.StatusBadRequest, u.Email)
	}
	if len(u.NIK) > 0 {
		if len(u.NIK) != 16 {
			return model.NewExpectedError("NIK must be 16 characters long", "USER_INVALID", http.StatusBadRequest, u.Email)
		}
		if match, _ := regexp.MatchString(numbersOnlyRegex, u.NIK); !match {
			return model.NewExpectedError("NIK must only consist of numbers", "USER_INVALID", http.StatusBadRequest, u.Email)
		}
	}
	if len(u.PhoneNo) > 0 {
		if match, _ := regexp.MatchString(phoneNoRegex, u.PhoneNo); !match {
			return model.NewExpectedError("phone_no must be a valid phone number", "USER_INVALID", http.StatusBadRequest, u.Email)
		}
	}

	return nil
}

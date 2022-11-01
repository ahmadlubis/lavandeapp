package entity

import "time"

type UserRole uint8

const (
	UserRoleResident UserRole = iota + 1
	UserRoleAdmin
)

type UserResidenceStatus uint8

const (
	UserResidenceStatusNotApplicable UserResidenceStatus = iota + 1
	UserResidenceStatusRenter
)

type User struct {
	ID              uint                `gorm:"column:id"`
	Name            string              `gorm:"column:name;default:null"`
	NIK             string              `gorm:"column:nik;default:null"`
	Email           string              `gorm:"column:email;default:null"`
	PhoneNo         string              `gorm:"column:phone_no;default:null"`
	Role            UserRole            `gorm:"column:role;default:null"`
	ResidenceStatus UserResidenceStatus `gorm:"column:residence_status;default:null"`
	Password        []byte              `gorm:"column:password;default:null"`
	CreatedAt       time.Time           `gorm:"column:created_at;"`
	UpdatedAt       time.Time           `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}

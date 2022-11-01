package entity

import (
	"fmt"
	"time"
)

type UserRole uint8

const (
	UserRoleResident UserRole = iota + 1
	UserRoleAdmin
)

func (s UserRole) String() string {
	switch s {
	case UserRoleResident:
		return "resident"
	case UserRoleAdmin:
		return "admin"
	default:
		return "invalid_user_role"
	}
}

func ParseUserRole(s string) (UserRole, error) {
	switch s {
	case "resident":
		return UserRoleResident, nil
	case "admin":
		return UserRoleAdmin, nil
	default:
		return 0, fmt.Errorf("invalid user role")
	}
}

type UserResidenceStatus uint8

const (
	UserResidenceStatusNotApplicable UserResidenceStatus = iota + 1
	UserResidenceStatusRenter
)

func (s UserResidenceStatus) String() string {
	switch s {
	case UserResidenceStatusNotApplicable:
		return "N/A"
	case UserResidenceStatusRenter:
		return "renter"
	default:
		return "invalid_user_residence_status"
	}
}

func ParseUserResidenceStatus(s string) (UserResidenceStatus, error) {
	switch s {
	case "N/A":
		return UserResidenceStatusNotApplicable, nil
	case "renter":
		return UserResidenceStatusRenter, nil
	default:
		return 0, fmt.Errorf("invalid user role")
	}
}

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

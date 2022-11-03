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

type UserStatus uint8

const (
	UserStatusActive UserStatus = iota + 1
	UserStatusNonactive
)

func (s UserStatus) String() string {
	switch s {
	case UserStatusActive:
		return "active"
	case UserStatusNonactive:
		return "nonactive"
	default:
		return "invalid_user_residence_status"
	}
}

func ParseUserStatus(s string) (UserStatus, error) {
	switch s {
	case "active":
		return UserStatusActive, nil
	case "nonactive":
		return UserStatusNonactive, nil
	default:
		return 0, fmt.Errorf("invalid user status")
	}
}

type User struct {
	ID        uint       `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name;default:null" json:"name"`
	NIK       string     `gorm:"column:nik;default:null" json:"nik"`
	Email     string     `gorm:"column:email;default:null" json:"email"`
	PhoneNo   string     `gorm:"column:phone_no;default:null" json:"phone_no"`
	Role      UserRole   `gorm:"column:role;default:null" json:"role"`
	Status    UserStatus `gorm:"column:status;default:null" json:"status"`
	Password  []byte     `gorm:"column:password;default:null" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at;" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

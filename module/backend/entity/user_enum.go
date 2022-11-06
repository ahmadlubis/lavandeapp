package entity

import (
	"encoding/json"
	"fmt"
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

func (s UserRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
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

func (s UserStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
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

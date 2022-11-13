package entity

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
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
		return 0, model.NewExpectedError("invalid user status", "USER_INVALID", http.StatusBadRequest, "")
	}
}

type UserReligion uint8

const (
	UserReligionOthers UserReligion = iota + 1
	UserReligionIslam
	UserReligionProtestant
	UserReligionCatholic
	UserReligionHindu
	UserReligionBuddha
	UserReligionConfucius
)

func (s UserReligion) String() string {
	switch s {
	case UserReligionOthers:
		return "others"
	case UserReligionIslam:
		return "islam"
	case UserReligionProtestant:
		return "protestant"
	case UserReligionCatholic:
		return "catholic"
	case UserReligionHindu:
		return "hindu"
	case UserReligionBuddha:
		return "buddha"
	case UserReligionConfucius:
		return "confucius"
	default:
		return "invalid_user_religion"
	}
}

func (s UserReligion) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func ParseUserReligion(s string) (UserReligion, error) {
	switch s {
	case "others":
		return UserReligionOthers, nil
	case "islam":
		return UserReligionIslam, nil
	case "protestant":
		return UserReligionProtestant, nil
	case "catholic":
		return UserReligionCatholic, nil
	case "hindu":
		return UserReligionHindu, nil
	case "buddha":
		return UserReligionBuddha, nil
	case "confucius":
		return UserReligionConfucius, nil
	default:
		return 0, model.NewExpectedError("invalid user religion", "USER_INVALID", http.StatusBadRequest, "")
	}
}

package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type TenantRole uint8

const (
	TenantRoleRenter TenantRole = iota + 1
	TenantRoleRenterFamily
)

func (s TenantRole) String() string {
	switch s {
	case TenantRoleRenter:
		return "renter"
	case TenantRoleRenterFamily:
		return "renter_family"
	default:
		return "invalid_tenant_role"
	}
}

func (s TenantRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func ParseTenantRole(s string) (TenantRole, error) {
	switch s {
	case "renter":
		return TenantRoleRenter, nil
	case "renter_family":
		return TenantRoleRenterFamily, nil
	default:
		return 0, fmt.Errorf("invalid tenant role")
	}
}

type Tenant struct {
	ID        uint64     `gorm:"column:id" json:"id"`
	UnitID    uint64     `gorm:"column:unit_id" json:"unit_id"`
	UserID    uint64     `gorm:"column:user_id" json:"user_id"`
	Role      TenantRole `gorm:"column:role" json:"role"`
	StartAt   *time.Time `gorm:"column:start_at" json:"start_at"`
	EndAt     *time.Time `gorm:"column:end_at" json:"end_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (Tenant) TableName() string {
	return "tenants"
}

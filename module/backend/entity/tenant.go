package entity

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
	"time"
)

type TenantRole uint8

const (
	TenantRoleOwner TenantRole = iota + 1
	TenantRoleRenter
	TenantRoleResident
	TenantRoleStaff
)

func (s TenantRole) String() string {
	switch s {
	case TenantRoleOwner:
		return "owner"
	case TenantRoleRenter:
		return "renter"
	case TenantRoleResident:
		return "resident"
	case TenantRoleStaff:
		return "staff"
	default:
		return "invalid_tenant_role"
	}
}

func (s TenantRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func ParseTenantRole(s string) (TenantRole, error) {
	switch s {
	case "owner":
		return TenantRoleOwner, nil
	case "renter":
		return TenantRoleRenter, nil
	case "resident":
		return TenantRoleResident, nil
	case "staff":
		return TenantRoleStaff, nil
	default:
		return 0, model.NewExpectedError("invalid tenant role", "TENANT_INVALID", http.StatusBadRequest, "")
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

func (t Tenant) Validate() error {
	if t.StartAt != nil && t.EndAt != nil {
		if !t.StartAt.Before(*t.EndAt) {
			return model.NewExpectedError("end_at must be after start_at", "TENANT_INVALID", http.StatusBadRequest, "")
		}
	}
	return nil
}

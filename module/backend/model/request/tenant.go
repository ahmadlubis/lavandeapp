package request

import "time"

type CreateTenantRequest struct {
	UnitID  uint64     `json:"unit_id"`
	UserID  uint64     `json:"user_id"`
	Role    string     `json:"role"`
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}

type DeleteTenantRequest struct {
	UnitID uint64 `json:"unit_id"`
	UserID uint64 `json:"user_id"`
}

type ListTenantRequest struct {
	UnitID     uint64 `json:"unit_id"`
	UserID     uint64 `json:"user_id"`
	ActiveOnly bool   `json:"active_only"`
	Limit      uint64 `json:"limit"`
	Offset     uint64 `json:"offset"`
}

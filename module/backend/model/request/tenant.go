package request

import "time"

type CreateTenantRequest struct {
	UnitID  uint64     `json:"unit_id"`
	UserID  uint64     `json:"user_id"`
	Role    string     `json:"role"`
	StartAt *time.Time `json:"start_at"`
	EndAt   *time.Time `json:"end_at"`
}

package response

import "github.com/ahmadlubis/lavandeapp/module/backend/entity"

type ListTenantResponse struct {
	Data []entity.Tenant `json:"data"`
	Meta PaginationMeta  `json:"meta"`
}

package response

import "github.com/ahmadlubis/lavandeapp/module/backend/entity"

type ListUnitResponse struct {
	Data []entity.Unit  `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

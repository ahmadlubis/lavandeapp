package response

import "github.com/ahmadlubis/lavandeapp/module/backend/entity"

type PaginationMeta struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
	Count  uint64 `json:"count"`
	Total  uint64 `json:"total"`
}

type ListUserResponse struct {
	Data []entity.User  `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

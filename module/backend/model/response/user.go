package response

import "github.com/ahmadlubis/lavandeapp/module/backend/entity"

type ListUserResponse struct {
	Data []entity.User  `json:"data"`
	Meta PaginationMeta `json:"meta"`
}

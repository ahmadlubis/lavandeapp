package response

type PaginationMeta struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
	Count  uint64 `json:"count"`
	Total  uint64 `json:"total"`
}

type GenericListResponse struct {
	Data []map[string]string `json:"data"`
	Meta PaginationMeta      `json:"meta"`
}

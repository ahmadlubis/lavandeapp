package request

type CreateUnitRequest struct {
	GovID  string `json:"gov_id"`
	Tower  string `json:"tower"`
	Floor  string `json:"floor"`
	UnitNo string `json:"unit_no"`
	AJB    string `json:"ajb"`
	Akte   string `json:"akte"`
}

type UpdateUnitRequest struct {
	ID   uint64 `json:"id"`
	AJB  string `json:"ajb"`
	Akte string `json:"akte"`
}

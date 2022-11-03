package request

type RegisterUserRequest struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

package request

type RegisterUserRequest struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_no"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	TargetEmail string `json:"target_email"`
	Name        string `json:"name"`
	NIK         string `json:"nik"`
	Email       string `json:"email"`
	PhoneNo     string `json:"phone_no"`
	Password    string `json:"password"`
}

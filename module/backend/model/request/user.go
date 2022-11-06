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

type SelfUpdateUserRequest struct {
	TargetEmail string `json:"-"`
	Name        string `json:"name"`
	NIK         string `json:"nik"`
	Email       string `json:"email"`
	PhoneNo     string `json:"phone_no"`
	Password    string `json:"password"`
}

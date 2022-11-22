package request

type RegisterUserRequest struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_no"`
	Religion string `json:"religion"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	TargetEmail string  `json:"target_email"`
	Name        *string `json:"name"`
	NIK         *string `json:"nik"`
	Email       *string `json:"email"`
	PhoneNo     *string `json:"phone_no"`
	Religion    *string `json:"religion"`
	Password    *string `json:"password"`
}

type AdminUpdateUserRequest struct {
	TargetId uint64  `json:"target_id"`
	Status   *string `json:"status"`
}

type ListUserRequest struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	Email    string `json:"email"`
	PhoneNo  string `json:"phone_no"`
	Status   string `json:"status"`
	Religion string `json:"religion"`
	Limit    uint64 `json:"limit"`
	Offset   uint64 `json:"offset"`
}

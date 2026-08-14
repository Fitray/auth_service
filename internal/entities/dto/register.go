package dto

type RegisterRequest struct {
	Login       string
	Password    string
	PhoneNumber *string
}

type RegisterResponse struct {
	UserId int64
}

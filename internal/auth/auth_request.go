package auth

type AuthRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	UserType    string `json:"userType" validate:"required"`
}

type UserReq struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
	UserType    string `json:"userType" validate:"required"`
}

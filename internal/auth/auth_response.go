package auth

type AuthResponse struct {
	Token       string `json:"paseto-token"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}

package auth

type AuthResponse struct {
	Token string `json:"token"`
	PhoneNumber    string `json:"phoneNumber"`
	Role    string `json:"role"`
}

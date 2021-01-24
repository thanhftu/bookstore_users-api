package users

// LoginRequest contains email and password request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

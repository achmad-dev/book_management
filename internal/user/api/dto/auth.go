package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

package userschema

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type SignUpResponse struct {
	Token    string `json:"token"`
	SigndURL string `json:"signed_url"`
}

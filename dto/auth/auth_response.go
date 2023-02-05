package authdto

type LoginResponse struct {
	Id       int    ` json:"id"`
	Username string ` json:"username"`
	Fullname string ` json:"fullname"`
	Token    string ` json:"token"`
}

type AuthResponse struct {
	Id       int    ` json:"id"`
	Username string ` json:"username"`
	Fullname string ` json:"fullname"`
}

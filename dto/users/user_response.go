package usersdto

type UserResponse struct {
	Id_user  int    `json:"id_user"`
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
}

package usersdto

type CreatUserRequest struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Status   string `json:"status" form:"status" validate:"required"`
}
type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Status   string `json:"status" form:"status"`
}

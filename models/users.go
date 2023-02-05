package models

type User struct {
	Id_User  int    `json:"id_user" gorm:"primary_key:auto_increment"`
	Fullname string `json:"fullname" gorm:"type:varchar(255)"`
	Username string `json:"username" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Status   string `json:"status" gorm:"type:varchar(100)"`
}

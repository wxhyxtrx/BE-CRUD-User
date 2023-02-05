package repositories

import (
	"ptedi/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(username string) (models.User, error)
	CekUser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}
func (r *repository) CekUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	return user, err
}

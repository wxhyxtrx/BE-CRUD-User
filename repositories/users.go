package repositories

import (
	"ptedi/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	FindAllUser() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	CheckUsername(username string) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) FindAllUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error
	return user, err
}

func (r *repository) CheckUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "username=?", username).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}

package repositories

import "testAPI/masters/apis/models"

type UserRepo interface {
	GetAllUsers()([]*models.Users, error)
	GetUserById(id int)(*models.Users,error)
	InsertUser(user *models.Users) error
	UpdateUser(id int, user *models.Users) error
}

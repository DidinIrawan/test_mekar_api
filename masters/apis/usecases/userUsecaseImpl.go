package usecases

import (
	"testAPI/masters/apis/models"
	"testAPI/masters/apis/repositories"
)

type UserUsecasesImpl struct {
	userRepo repositories.UserRepo
}

func (u UserUsecasesImpl) UpdateUser(id int, user *models.Users) error {
	err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecasesImpl) InsertUser(user *models.Users) error {
	err := u.userRepo.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecasesImpl) GetUserById(id int) (*models.Users, error) {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecasesImpl) GetAllUsers() ([]*models.Users, error) {
	jobs, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return jobs, nil
}

func InitUserUsecase(userRepo repositories.UserRepo) UserUsecases {
	return &UserUsecasesImpl{userRepo}
}
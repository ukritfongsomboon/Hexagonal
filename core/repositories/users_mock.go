package repositories

import (
	"errors"
	"hexagonal/core/models"
)

type userRepositoryMock struct {
	users []models.UserModel
}

func NewUserRepositoryMock() userRepositoryMock {
	users := []models.UserModel{
		{UserID: "136e4d95-41a7-4d50-9c9d-e25f93fa406a", Email: "kobori4268@gmail.com", Password: "4813494d137e1631bba301d5acab6e7bb7aa74ce1185d456565ef51d737677b2", Name: "user1", Role: 0, Status: false},
		{UserID: "136e4d95-41a7-4d50-9c9d-e25f93fa406b", Email: "kobori4268@gmail.com", Password: "4813494d137e1631bba301d5acab6e7bb7aa74ce1185d456565ef51d737677b2", Name: "user2", Role: 0, Status: false},
	}
	return userRepositoryMock{users: users}
}

func (r userRepositoryMock) GetAll(p models.UserPaginationModel) ([]models.UserModel, error) {
	return r.users, nil
}

func (r userRepositoryMock) GetById(userid string) (*models.UserModel, error) {
	for _, user := range r.users {
		if userid == user.UserID {
			return &user, nil
		}
	}
	return nil, errors.New("mongo: no documents in result")
}

func (r userRepositoryMock) Create(payload models.UserCreateModel) (*models.UserModel, error) {
	return nil, nil
}

func (r userRepositoryMock) CountAll() (int64, error) {
	return 0, nil
}

func (r userRepositoryMock) GetByEmail(email string) (*models.UserModel, error) {
	return nil, nil
}

package repositories

import "hexagonal/core/models"

type UserRepository interface {
	GetAll(models.UserPaginationModel) ([]models.UserModel, error)
	GetById(string) (*models.UserModel, error)
	GetByEmail(string) (*models.UserModel, error)
	GetCountAll() (int64, error)
	Create(models.UserCreateModel) (*models.UserModel, error)
	Update(models.UserUpdateReqModel) (*models.UserModel, error)
}

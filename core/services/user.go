package services

import "hexagonal/core/models"

type UserService interface {
	GetUsers(models.UserPaginationModel) (*models.UserResGetAllModel, error)
	GetUser(string) (*models.UserResModel, error)
	UpdateUser(models.UserUpdateReqModel) error
	UpdateUserImage(string, string, []byte) error
	SignIn(*models.SignInReqModel) (*models.SignInResModel, error)
	SignUp(*models.SignUpReqModel) (*models.SignUpResModel, error)
}

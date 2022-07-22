package services

import "hexagonal/core/models"

type MenuService interface {
	GetMenus() ([]models.MenuResModel, error)
	CreateMenu(models.MenuCreateReqModel) (*models.MenuCreateResModel, error)
	UpdateMenu() error
	DeleteMenu() error
}

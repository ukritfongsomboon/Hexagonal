package repositories

import "hexagonal/core/models"

type MenuRepository interface {
	GetAllMenu() ([]models.MenuModel, error)
	GetMenuByPer() error
	CreateMenu(models.MenuCreateReqModel) (*models.MenuModel, error)
	UpdateMenu() error
	DeleteMenu() error
}

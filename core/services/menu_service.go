package services

import (
	"hexagonal/common/logs"
	"hexagonal/core/models"
	"hexagonal/core/repositories"

	"hexagonal/utils"
)

type menuService struct {
	log      logs.AppLog
	menuRepo repositories.MenuRepository
}

func NewMenuService(log logs.AppLog, menuRepo repositories.MenuRepository) MenuService {
	return menuService{log: log, menuRepo: menuRepo}
}

func (s menuService) GetMenus() ([]models.MenuResModel, error) {
	menus, err := s.menuRepo.GetAllMenu()
	if err != nil {
		s.log.Error(err)
		return nil, utils.HandlerError{
			Code:    500,
			Message: "unexpected error",
		}
	}

	// # DTO Data Tranfer Object
	var menuResposne []models.MenuResModel
	for _, menu := range menus {
		menuItem := models.MenuResModel{
			Id:          menu.Id,
			Name:        menu.Name,
			Icon:        menu.Icon,
			Path:        menu.Path,
			PathName:    menu.PathName,
			Description: menu.Description,
		}
		menuResposne = append(menuResposne, menuItem)
	}
	return menuResposne, nil
}

func (s menuService) CreateMenu(m models.MenuCreateReqModel) (*models.MenuCreateResModel, error) {

	// # Data tranfer object handeler --> service
	reqData := models.MenuCreateReqModel{
		Role:        m.Role,
		Name:        m.Name,
		Icon:        m.Icon,
		Path:        m.Path,
		PathName:    m.PathName,
		Description: m.Description,
	}

	newMenu, err := s.menuRepo.CreateMenu(reqData)
	if err != nil {
		s.log.Error(err)
		return nil, utils.HandlerError{
			Code:    500,
			Message: "unexpected error",
		}
	}

	// # Data tranfer object Service --> Handler
	resData := models.MenuCreateResModel{
		Name:        newMenu.Name,
		Icon:        newMenu.Icon,
		Path:        newMenu.Path,
		PathName:    newMenu.PathName,
		Description: newMenu.Description,
	}

	return &resData, nil
}
func (s menuService) UpdateMenu() error {
	return nil
}
func (s menuService) DeleteMenu() error {
	return nil
}

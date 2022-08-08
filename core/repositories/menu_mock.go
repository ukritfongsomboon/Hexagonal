package repositories

import (
	"fmt"
	"hexagonal/core/models"
	"time"

	"github.com/google/uuid"
)

type menuRepositoryMock struct {
	menu []models.MenuModel
}

func NewMenuRepositoryMock() MenuRepository {
	menu := []models.MenuModel{}
	for i := 1; i <= 2; i++ {
		menu = append(menu, models.MenuModel{
			Id:          uuid.New().String(),
			Role:        1,
			Name:        fmt.Sprintf("name-%v", i),
			Icon:        "testicon",
			Path:        fmt.Sprintf("path-%v", i),
			PathName:    models.PathName{Name: fmt.Sprintf("name-%v", i)},
			Description: fmt.Sprintf("description-%v", i),
			Index:       i,
			Status:      true,
			CreatedDate: time.Now(),
			LastUpdate:  time.Now(),
		})
	}
	return &menuRepositoryMock{menu: menu}
}

func (r *menuRepositoryMock) GetAllMenu() ([]models.MenuModel, error) {
	return r.menu, nil
}

func (r *menuRepositoryMock) GetMenuByPer() error {
	return nil
}

func (r *menuRepositoryMock) CreateMenu(menu models.MenuCreateReqModel) (*models.MenuModel, error) {
	newMenu := models.MenuModel{
		Id:          uuid.New().String(),
		Role:        menu.Role,
		Name:        menu.Name,
		Icon:        menu.Icon,
		Path:        menu.Path,
		PathName:    menu.PathName,
		Description: menu.Description,
		Index:       0,
		Status:      false,
		CreatedDate: time.Now(),
		LastUpdate:  time.Now(),
	}

	r.menu = append(r.menu, newMenu)
	return &newMenu, nil
}

func (r *menuRepositoryMock) UpdateMenu() error {
	return nil
}

func (r *menuRepositoryMock) DeleteMenu() error {
	return nil
}

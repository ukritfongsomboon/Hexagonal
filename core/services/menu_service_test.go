package services_test

import (
	"fmt"
	"hexagonal/common/logs"
	"hexagonal/core/models"
	"hexagonal/core/repositories"
	"hexagonal/core/services"
	"reflect"
	"testing"
)

func Test_menuService_GetMenus(t *testing.T) {
	log := logs.NewAppLogs()
	menuRepository := repositories.NewMenuRepositoryMock()
	menuServices := services.NewMenuService(log, menuRepository)

	type fields struct {
		log      logs.AppLog
		menuRepo repositories.MenuRepository
	}

	tests := []struct {
		name    string
		fields  fields
		args    int
		want    []models.MenuResModel
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Case 1"},
		{name: "Case 2"},
		{name: "Case 3"},
		{name: "Case 4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := menuServices
			got, err := s.GetMenus()
			fmt.Println(err != nil)
			// TODO เมื่อได้รับ Error คืนกลับมาจาก Service
			if (err != nil) != tt.wantErr {
				t.Errorf("menuService.GetMenus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO เมื่อไม่มี Error จะทำการ Check Response จาก Service
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuService.GetMenus() = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_menuService_CreateMenu(t *testing.T) {
	log := logs.NewAppLogs()
	menuRepository := repositories.NewMenuRepositoryMock()
	menuServices := services.NewMenuService(log, menuRepository)

	type fields struct {
		log      logs.AppLog
		menuRepo repositories.MenuRepository
	}

	tests := []struct {
		name    string
		fields  fields
		args    models.MenuCreateReqModel
		want    *models.MenuCreateResModel
		wantErr bool
	}{
		{
			name: "case 1",
			args: models.MenuCreateReqModel{
				Role: 1,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := menuServices
			got, err := s.CreateMenu(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("menuService.CreateMenu() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("menuService.CreateMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}

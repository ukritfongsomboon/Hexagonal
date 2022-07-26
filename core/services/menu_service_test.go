package services_test

import (
	"fmt"
	"hexagonal/common/logs"
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
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Case 1", want: 5},
		{name: "Case 2", want: 5},
		{name: "Case 3", want: 5},
		{name: "Case 4", want: 5},
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
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("menuService.GetMenus() = %v, want %v", len(got), tt.want)
			}

		})
	}
}

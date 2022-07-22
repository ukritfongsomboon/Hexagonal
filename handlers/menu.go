package handlers

import (
	"hexagonal/core/models"
	"hexagonal/core/services"
	"hexagonal/utils"

	"github.com/gofiber/fiber/v2"
)

type menuHandler struct {
	menuSrv services.MenuService
}

func NewMenuHandler(menuSrv services.MenuService) menuHandler {
	return menuHandler{menuSrv: menuSrv}
}

func (h menuHandler) GetMenu(c *fiber.Ctx) error {
	menus, err := h.menuSrv.GetMenus()
	if err != nil {
		appErr, ok := err.(utils.HandlerError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"code":    appErr.Code,
				"status":  false,
				"message": appErr.Message,
				"data":    "",
			})
		}
	}

	// FIX null to []
	if menus == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"status":  true,
			"message": "get menu success",
			"data":    make([]int, 0),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "get menu success",
		"data":    menus,
	})
}

func (h menuHandler) CreateMenu(c *fiber.Ctx) error {
	body := new(models.MenuCreateReqModel)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to parse body",
			"data":    "",
		})
	}
	menu, err := h.menuSrv.CreateMenu(*body)
	if err != nil {
		appErr, ok := err.(utils.HandlerError)
		if ok {
			return c.Status(appErr.Code).JSON(fiber.Map{
				"code":    appErr.Code,
				"status":  false,
				"message": appErr.Message,
				"data":    "",
			})
		}
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  true,
		"message": "create menu success",
		"data":    menu,
	})
}

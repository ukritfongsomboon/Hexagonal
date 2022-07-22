package handlers

import (
	"hexagonal/core/models"
	"hexagonal/core/services"
	"hexagonal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

type JSONResult struct {
	Code    int         `json:"code" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} //@name Response

// @Summary      Get user
// @Description  get all user in services
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        row  query     int     false  "row of page"
// @Param        page query     int     false  "page"
// @Success      200  {object}  models.UserResGetAllModel
// @success      200 {object}  JSONResult{data=[]string} "desc"
// @Failure      400  {object}  models.UserResGetAllModel
// @Failure      404  {object}  models.UserResGetAllModel
// @Failure      500  {object}  models.UserResGetAllModel
// @Router        /api/v1/users [get]
// @Security     Authorization
func (h userHandler) GetUsers(c *fiber.Ctx) error {
	var p models.UserPaginationModel
	p.Page, _ = strconv.Atoi(c.Query("page", "1"))
	p.Row, _ = strconv.Atoi(c.Query("row", "10"))

	users, err := h.userSrv.GetUsers(p)
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
	if users == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code":    200,
			"status":  true,
			"message": "get user success",
			"data":    make([]int, 0),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "get user success",
		"data":    users,
	})
}

// @Summary      Show an account
// @Description  get string by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  models.UserResGetAllModel
// @success      200 {object}  JSONResult{data=[]string} "desc"
// @Failure      400  {object}  models.UserResGetAllModel
// @Failure      404  {object}  models.UserResGetAllModel
// @Failure      500  {object}  models.UserResGetAllModel
// @Router       /api/v1/user/{id}/account [get]
func (h userHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("userid")
	users, err := h.userSrv.GetUser(userID)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    200,
		"status":  true,
		"message": "get user success",
		"data":    users,
	})
}

// @Tags         Sign In
// @Summary      Sign In with with username and password
// @Description  Get accesstoken with username and password
// @Router       /api/v1/signin [POST]
// @Accept       json
// @Produce      json
// @Param        Body   body    models.SignInReqModel  true "username & password"
// @Success      200  {object}  JSONResult
// @Failure      400  {object}  JSONResult
// @Failure      404  {object}  JSONResult
// @Failure      500  {object}  JSONResult
func (h userHandler) SignIn(c *fiber.Ctx) error {
	body := new(models.SignInReqModel)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to parse body",
			"data":    "",
		})
	}

	data, err := h.userSrv.SignIn(body)
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

	// clear cookie client
	c.ClearCookie()

	// create cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "Accesstoken"
	cookie.Value = data.Accesstoken
	cookie.Secure = false
	cookie.SessionOnly = true
	cookie.MaxAge = 3000
	cookie.Expires = time.Now().Add(10 * time.Second)

	// set cookie
	c.Cookie(cookie)

	// # Success Case
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"status":  true,
		"message": "login success",
		"data":    data,
	})
}

func (h userHandler) SignUp(c *fiber.Ctx) error {
	body := new(models.SignUpReqModel)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"status":  false,
			"message": "Failed to parse body",
			"data":    "",
		})
	}

	user, err := h.userSrv.SignUp(body)

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

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"status":  true,
		"message": "create user success",
		"data":    user,
	})
}

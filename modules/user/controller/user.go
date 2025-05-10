package controller

import "backend/modules/user/service"

type (
	UserController interface {
		// DefaultController
	}

	userController struct {
		userService service.UserService
	}
)

// declare NewUserController function
func NewUserController(userService service.UserService) UserController {
	return &userController{userService: userService}
}

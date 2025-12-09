package controller

import (
	"GolangBackendEcommerce/internal/service"
	"GolangBackendEcommerce/pkg/response"

	"github.com/gin-gonic/gin"
)

// type UserController struct {
// 	userService *service.UserService
// }

// // use this to make our routers acknowledge which controller we are calling
// func NewUserController() *UserController {
// 	return &UserController{
// 		userService: service.NewUserService(),
// 	}
// }

// func (uc *UserController) GetUserById(c *gin.Context) {
// 	// if err != nil {
// 	// 	response.ErrorResponse(c, 20003, "Invalid parameter")
// 	// 	return
// 	// }
// 	// response.SuccessResponse(c, 20001, []string{"User1", "User2"})

// }

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (uc *userController) Register(c *gin.Context) {
	result := uc.userService.Register("email", "purpose")
	response.SuccessResponse(c, result, nil)
}

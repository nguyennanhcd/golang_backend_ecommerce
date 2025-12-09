package user

import (
	"GolangBackendEcommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routes
	// this is non-dependency
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandleNoneDependency := controller.NewUserController(us)

	// wire
	userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// use wire go to handle dependency injection later

	// private routes
	userRouterPrivate := Router.Group("/user")
	// userRouterPrivate.Use(Limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("get_info")
	}
}

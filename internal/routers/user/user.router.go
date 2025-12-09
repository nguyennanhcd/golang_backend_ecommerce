package user

import (
	"GolangBackendEcommerce/internal/controller"
	"GolangBackendEcommerce/internal/repo"
	"GolangBackendEcommerce/internal/service"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public routes
	// this is non-dependency
	ur := repo.NewUserRepository()
	us := service.NewUserService(ur)
	userHandleNoneDependency := controller.NewUserController(us)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userHandleNoneDependency.Register)
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

package routers

import (
	"GolangBackendEcommerce/internal/routers/manager"
	"GolangBackendEcommerce/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)

package global

import (
	"GolangBackendEcommerce/pkg/logger"
	"GolangBackendEcommerce/pkg/settings"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)

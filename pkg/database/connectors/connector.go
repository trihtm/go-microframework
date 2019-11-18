package connectors

import (
	"github.com/jinzhu/gorm"
	"go-microframework/internal/config"
)

type Connector interface {
	Connect(config config.DatabaseConnectionConfiguration) *gorm.DB
}
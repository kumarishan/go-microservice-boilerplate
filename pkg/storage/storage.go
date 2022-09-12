package storage

import (
	"github.com/kumarishan/go-microservice-boilerplate/pkg/config"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _ = di.Provide(NewGormDB)

func NewGormDB(config *config.Config, log *logger.Logger) (*gorm.DB, error) {
	dsn := "root:test@tcp(127.0.0.1:3306)/catalog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

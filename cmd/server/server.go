package server

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/internal/products"
	"github.com/kumarishan/go-microservice-boilerplate/internal/rest"
	_ "github.com/kumarishan/go-microservice-boilerplate/internal/rest/products"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/config"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
	_ "github.com/kumarishan/go-microservice-boilerplate/pkg/storage"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

type server struct {
	router *gin.Engine
	logger *logger.Logger
}

func NewServer(e *gin.Engine) *server {

	di.Provide(func() *gin.Engine { return e })
	di.Provide(logger.NewLogger)
	di.Provide(config.NewConfig)

	return &server{
		router: e,
	}
}

func (s *server) Start() error {

	// register routes
	if err := di.Invoke(func(r struct {
		dig.In
		R rest.Routes `name:"api/products"`
	}) {
		r.R.SetupRoutes()
	}); err != nil {
		fmt.Println(err)
	}

	// Db Migrate
	if err := di.Invoke(func(r struct {
		dig.In
		Db  *gorm.DB
		Log *logger.Logger
	}) {
		if err := products.MigrateDb(r.Db); err != nil {
			r.Log.Errorw(context.Background(), "error while migrating db", err)
		}
	}); err != nil {
		fmt.Printf("error while migrating %v\n", err)
	}

	// start the server
	return s.router.Run(fmt.Sprintf(":%s", "8080"))
}

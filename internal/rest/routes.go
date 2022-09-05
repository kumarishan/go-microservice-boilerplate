package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/di"
	"go.uber.org/dig"
)

type Routes interface {
	SetupRoutes()
}

var _ = di.Provide(NewAPIV1Routes, dig.Name("api"))

func NewAPIV1Routes(g *gin.Engine) *gin.RouterGroup {
	return g.Group("api/v1")
}

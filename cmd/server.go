package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/pkg/logger"
)

type server struct {
	router *gin.Engine
	logger *logger.Logger
}

func NewServer(e *gin.Engine) *server {

	return &server{
		router: e,
	}
}

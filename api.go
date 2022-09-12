package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kumarishan/go-microservice-boilerplate/cmd/server"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {
	g := gin.Default()
	srv := server.NewServer(g)
	return srv.Start()
}

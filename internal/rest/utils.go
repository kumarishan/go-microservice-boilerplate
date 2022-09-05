package rest

import "github.com/gin-gonic/gin"

func GinHandler[D any](handler func(*gin.Context) (int, D, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code, data, err := handler(ctx)
		if err != nil {
			ctx.JSON(code, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(code, data)
		}
	}
}

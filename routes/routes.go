package routes

import (
	"sprintone/services/calculator"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.POST("/api/v1/calculate", func(ctx *gin.Context) {
		calculator.CalculatorGet(ctx)
	})
}

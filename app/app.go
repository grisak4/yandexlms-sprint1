package app

import (
	"sprintone/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	r := gin.Default()

	routes.InitRoutes(r)

	r.Run(":8081")
}

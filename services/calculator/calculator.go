package calculator

import (
	"log"
	"sprintone/calc"

	"github.com/gin-gonic/gin"
)

func CalculatorGet(c *gin.Context) {
	exp := struct {
		Expression string `json:"expression"`
	}{}

	if err := c.ShouldBindJSON(&exp); err != nil {
		log.Printf("[ERROR] %s", err.Error())
		c.JSON(422, gin.H{
			"error": "Expression is not valid",
		})
		return
	}
	if exp.Expression == "" {
		log.Printf("[ERROR] Expression is empty")
		c.JSON(422, gin.H{
			"error": "Expression is not valid",
		})
		return
	}

	result, err := calc.Calculator(exp.Expression)
	if err != nil {
		log.Printf("[ERROR] %s", err.Error())
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": result,
	})
}

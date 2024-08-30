package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()
	//definindo a rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//router rodando Api
	router.Run(":8080")
}

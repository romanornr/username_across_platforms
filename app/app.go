package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

func StartApp() {
	route()

	port := 8080

	router.Run(":" + port)
}

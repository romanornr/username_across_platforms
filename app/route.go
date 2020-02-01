package app

import (
	"github.com/romanornr/username_across_platforms/controller"
	"github.com/romanornr/username_across_platforms/middleware"
)

func route() {
	router.Use(middleware.CORSMiddleware())

	router.POST("/username", controller.Username)
}

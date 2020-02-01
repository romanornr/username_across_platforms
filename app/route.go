package app

import "github.com/romanornr/username_across_platforms/controller"

func route() {
	router.Use(middleware.CORSMiddleware())

	router.POST("/username", controller.Username())
}

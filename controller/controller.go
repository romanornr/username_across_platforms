package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/romanornr/username_across_platforms/service"
	"net/http"
)

func Username(c *gin.Context) {
	var urls []string
	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid json body"))
		return
	}
	matchedUrls := service.UsernameService.UsernameCheck(urls)

	c.JSON(http.StatusOK, matchedUrls)
}

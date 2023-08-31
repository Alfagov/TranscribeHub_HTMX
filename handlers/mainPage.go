package handlers

import (
	"github.com/gin-gonic/gin"
)

func MainPageHandler(c *gin.Context) {
	auth := c.GetBool("logged_in")
	c.HTML(200, "mainPage.html", gin.H{
		"IsAuthenticated": auth,
	})
}

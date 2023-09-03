package handlers

import (
	"github.com/gin-gonic/gin"
)

func MainPageHandler(c *gin.Context) {
	auth := c.GetBool("logged_in")

	c.HTML(200, "pages/mainPage.html", gin.H{
		"IsAuthenticated": auth,
		"Counters": []struct {
			Text          string
			Value         int
			Max           int
			ProgressClass string
		}{
			{Text: "Transcriptions",
				Value:         7,
				Max:           10,
				ProgressClass: "progress-secondary",
			},
			{
				Text:          "Transcriptions",
				Value:         3,
				Max:           10,
				ProgressClass: "progress-error",
			},
			{
				Text:          "Transcriptions",
				Value:         40,
				ProgressClass: "progress-info",
			},
		},
	})
}

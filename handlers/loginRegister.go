package handlers

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/models"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"log"
)

func LoginUserHandler(db database.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.LoginUser
		err := c.ShouldBind(&user)
		if err != nil {
			log.Println("LOGIN ERROR", err)
			c.HTML(200, "", gin.H{})
			return
		}

		dbUser, err := db.LoginUser(user)
		if err != nil {
			c.HTML(200, "", gin.H{})
			log.Println("LOGIN ERROR", err)
			return
		}

		c.Header("HX-Trigger", "loggedInEvent")
		c.SetCookie("Authentication", dbUser.Id, 3600, "/", "localhost", false, true)
		c.HTML(200, "sidebar-inner", gin.H{
			"IsAuthenticated": true,
			"NotPremium":      true,
		})

		log.Println("LOGIN USER", dbUser)
	}
}

func LogoutUserHandler(db database.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {

		c.SetCookie("Authentication", "", -1, "/", "localhost", false, true)
		c.HTML(200, "sidebar-header", gin.H{
			"IsAuthenticated": false,
		})

	}
}

func RegisterUserHandler(db database.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.RegisterUser
		err := c.ShouldBind(&user)
		if err != nil {
			c.HTML(200, "", gin.H{})
			return
		}

		user.Id = ulid.Make().String()
		log.Println(user.Id)

		err = db.RegisterUser(user)
		if err != nil {
			c.HTML(200, "", gin.H{})
			return
		}
		c.HTML(200, "registeredModal.html", gin.H{})
	}
}

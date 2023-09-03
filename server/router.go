package server

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/handlers"
	"TranscribeHub_HTMX/middlewares"
	"TranscribeHub_HTMX/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newRouter(db database.Dao) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.LoadHTMLGlob("./templates/**/*")

	addStaticRoutes(router)
	addPageRoutes(router)
	addUserRoutes(router, db)
	addValidationRoutes(router, db)
	addModalsRoutes(router, db)

	return router
}

func addStaticRoutes(engine *gin.Engine) {
	engine.GET("/static/css", func(c *gin.Context) {
		c.File("./static/css.css")
	})
}

func addPageRoutes(engine *gin.Engine) {
	engine.GET("/", middlewares.AuthMiddleware(), handlers.MainPageHandler)
	engine.GET("/sidebar", middlewares.AuthMiddleware(), func(c *gin.Context) {
		c.HTML(200, "sidebar-inner", gin.H{
			"IsAuthenticated": c.GetBool("logged_in"),
		})
	})
	engine.GET("/home", middlewares.AuthMiddleware(), func(c *gin.Context) {
		if c.GetBool("logged_in") {
			c.HTML(200, "grid-page", gin.H{
				"IsAuthenticated": c.GetBool("logged_in"),
				"Counters": []struct {
					Text  string
					Value int
					Max   int
				}{
					{Text: "Transcriptions",
						Value: 7,
						Max:   10},
					{
						Text:  "Transcriptions",
						Value: 3,
						Max:   10,
					},
					{
						Text:  "Transcriptions",
						Value: 40,
					},
				},
			})
			return
		}
		c.HTML(200, "landing-page", gin.H{
			"IsAuthenticated": c.GetBool("logged_in"),
			"Counters": []struct {
				Text  string
				Value int
				Max   int
			}{
				{Text: "Transcriptions",
					Value: 7,
					Max:   10},
				{
					Text:  "Transcriptions",
					Value: 3,
					Max:   10,
				},
				{
					Text:  "Transcriptions",
					Value: 40,
				},
			},
		})
	})
}

func addUserRoutes(engine *gin.Engine, db database.Dao) {
	engine.POST("/user/login", handlers.LoginUserHandler(db))
	engine.POST("/user/register", handlers.RegisterUserHandler(db))
	engine.GET("/user/logout", handlers.LogoutUserHandler(db))

}

func addModalsRoutes(engine *gin.Engine, db database.Dao) {
	modals := engine.Group("/modals")
	{
		modals.GET("/login", func(c *gin.Context) {
			c.HTML(200, "general-lr-modal", gin.H{
				"ModalField": models.LoginFields,
				"loginType":  true,
				"Boost":      true,
				"FormUri":    "/user/login",
				"ModalId":    "login-modal",
			})
		})

		modals.GET("/register", func(c *gin.Context) {
			c.HTML(200, "general-lr-modal", gin.H{
				"ModalField":   models.RegisterFields,
				"registerType": true,
				"FormUri":      "/user/register",
				"ModalId":      "login-modal",
			})
		})
	}
}

func addValidationRoutes(engine *gin.Engine, db database.Dao) {
	engine.POST("/validate/email/:fieldId", handlers.ValidateEmailHandler())
	engine.POST("/validate/password/:fieldId", handlers.ValidatePasswordHandler())
	engine.POST("/validate/username/:fieldId", handlers.ValidateUsernameHandler(db))
}

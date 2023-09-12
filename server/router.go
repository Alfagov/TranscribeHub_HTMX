package server

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/handlers"
	"TranscribeHub_HTMX/middlewares"
	"TranscribeHub_HTMX/models"
	"TranscribeHub_HTMX/pageElements"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

func newRouter(db database.Dao) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	router.LoadHTMLGlob("./templates/**/*")

	temp, err := template.ParseGlob("./templates/**/*")
	if err != nil {
		panic(err)
	}

	tCreator := &pageElements.TemplateCreator{Template: temp}

	addStaticRoutes(router)
	addPageRoutes(router, tCreator, db)
	addUserRoutes(router, db)
	addValidationRoutes(router, db)
	addModalsRoutes(router, db)

	return router
}

func addStaticRoutes(engine *gin.Engine) {
	engine.GET("/static/css", func(c *gin.Context) {
		c.File("./static/css.css")
	})
	engine.GET("/static/sidebarjs", func(c *gin.Context) {
		c.File("./static/js/sidebarScript.js")
	})
}

func addPageRoutes(engine *gin.Engine, templates *pageElements.TemplateCreator, db database.Dao) {
	engine.GET("/", middlewares.AuthMiddleware(), handlers.MainPHandler(templates, db))
	engine.GET("/content", middlewares.AuthMiddleware(), handlers.HomeComponentHandler(db, templates))
	engine.GET("/dashboard", middlewares.AuthMiddleware(), handlers.DashBoardPage(db, templates))
	engine.GET("/pricing", middlewares.AuthMiddleware(), handlers.PricingPage(db, templates))
	engine.GET("/test", middlewares.AuthMiddleware(), func(c *gin.Context) {
		subCounters, err := db.GetDefaultCountersByUserId(c.GetString("user_id"))
		if err != nil {
			log.Println(err)
			return
		}

		nwElem, err := pageElements.GetListNewsAndConvert(db, templates)
		if err != nil {
			log.Println(err)
			return
		}
		c.HTML(200, "test-page", gin.H{
			"IsAuthenticated": true,
			"NotPremium":      true,
			"Counters":        subCounters.Counters,
			"News":            nwElem,
			"PricingCards":    models.PricingCards,
		})
	})
	sidebar := engine.Group("/sidebar")
	{
		sidebar.GET("/dashboard", func(c *gin.Context) {
			c.HTML(200, "sidebar-dashboard", gin.H{
				"IsAuthenticated": true,
				"NotPremium":      true,
			})
		})
	}
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
		modals.GET("/plans", func(c *gin.Context) {

		})
	}
}

func addValidationRoutes(engine *gin.Engine, db database.Dao) {
	engine.POST("/validate/email/:fieldId", handlers.ValidateEmailHandler())
	engine.POST("/validate/password/:fieldId", handlers.ValidatePasswordHandler())
	engine.POST("/validate/username/:fieldId", handlers.ValidateUsernameHandler(db))
}

package handlers

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/models"
	"TranscribeHub_HTMX/pageElements"
	"github.com/gin-gonic/gin"
)

func PricingPage(db database.Dao, templates *pageElements.TemplateCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetBool("logged_in") {
			c.Redirect(302, "/dashboard")
			return
		}

		outValues := pageElements.NewTemplateValuesMap()
		outValues = outValues.
			AddPricingCards(models.PricingCards).
			AddSidebarActiveElement("PlansActive").
			AddAuthentication(false)

		c.Header("HX-Push-Url", "/pricing")
		c.HTML(200, "pages/pricingMainPage.html", outValues)
	}
}

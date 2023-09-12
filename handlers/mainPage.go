package handlers

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/models"
	"TranscribeHub_HTMX/pageElements"
	"github.com/gin-gonic/gin"
	"log"
)

func MainPHandler(templates *pageElements.TemplateCreator, db database.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetBool("logged_in")
		outValues := pageElements.
			NewTemplateValuesMap().
			AddAuthentication(auth).
			AddSidebarActiveElement("HomeActive")

		if auth {
			nwElem, err := pageElements.GetListNewsAndConvert(db, templates)
			if err != nil {
				log.Println(err)
				return
			}

			subCounters, err := db.GetDefaultCountersByUserId(c.GetString("user_id"))
			if err != nil {
				log.Println(err)
				return
			}

			outValues = outValues.
				AddNewsList(nwElem).
				AddCounters(subCounters.Counters)

			c.Redirect(302, "/dashboard")

		} else {
			outValues = outValues.
				AddPricingCards(models.PricingCards)

			c.Header("HX-Push-Url", "/")
		}

		c.HTML(200, "pages/mainPage.html", outValues)

	}
}

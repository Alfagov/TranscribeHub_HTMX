package handlers

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/pageElements"
	"github.com/gin-gonic/gin"
	"log"
)

func DashBoardPage(db database.Dao, templates *pageElements.TemplateCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		outValues := pageElements.NewTemplateValuesMap()

		if !c.GetBool("logged_in") {
			c.Redirect(302, "/")
			return
		}

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
			AddAuthentication(true).
			AddNewsList(nwElem).
			AddCounters(subCounters.Counters)

		c.HTML(200, "pages/mainPage.html", outValues)
	}
}

func HomeComponentHandler(db database.Dao, templates *pageElements.TemplateCreator) gin.HandlerFunc {
	return func(c *gin.Context) {
		outValues := pageElements.NewTemplateValuesMap()

		if c.GetBool("logged_in") {

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
				AddAuthentication(true).
				AddNewsList(nwElem).
				AddCounters(subCounters.Counters)

			c.HTML(200, "grid-page-ui", outValues)
			return
		}
		c.HTML(200, "landing-page", outValues)
	}
}

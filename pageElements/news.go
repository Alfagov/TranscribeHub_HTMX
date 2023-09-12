package pageElements

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/models"
	"bytes"
	"html/template"
)

type TemplateCreator struct {
	*template.Template
}

func (tc *TemplateCreator) CreateNewsElement(news *models.News) (*template.HTML, error) {
	var outElement bytes.Buffer
	err := tc.ExecuteTemplate(&outElement, "news-element", news)
	if err != nil {
		return nil, err
	}
	outTemplate := template.HTML(outElement.String())
	return &outTemplate, nil
}

func (tc *TemplateCreator) CreateNewsList(newsList []*models.News) ([]*template.HTML, error) {
	var outList []*template.HTML
	for _, news := range newsList {
		tmpElement, err := tc.CreateNewsElement(news)
		if err != nil {
			return nil, err
		}

		outList = append(outList, tmpElement)
	}

	return outList, nil
}

func GetListNewsAndConvert(db database.Dao, tc *TemplateCreator) ([]*template.HTML, error) {
	news, err := db.GetNews()
	if err != nil {
		return nil, err
	}

	return tc.CreateNewsList(news)
}

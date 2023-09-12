package pageElements

import (
	"TranscribeHub_HTMX/models"
	"html/template"
)

type TemplateValuesMap map[string]any

func NewTemplateValuesMap() TemplateValuesMap {
	return make(TemplateValuesMap)
}

func (vm TemplateValuesMap) AddNewsList(news []*template.HTML) TemplateValuesMap {
	vm["News"] = news
	return vm
}

func (vm TemplateValuesMap) AddCounters(counters []*models.Counter) TemplateValuesMap {
	vm["Counters"] = counters
	return vm
}

func (vm TemplateValuesMap) AddAuthentication(auth bool) TemplateValuesMap {
	vm["IsAuthenticated"] = auth
	return vm
}

func (vm TemplateValuesMap) AddPricingCards(cards []models.PricingCard) TemplateValuesMap {
	vm["PricingCards"] = cards
	return vm
}

func (vm TemplateValuesMap) AddSidebarActiveElement(elementName string) TemplateValuesMap {
	vm[elementName] = "active"
	return vm
}

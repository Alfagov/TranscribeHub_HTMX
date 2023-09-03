package models

import "html/template"

type News struct {
	SvgNotification template.HTML
	Title           string
	Date            string
	Text            string
	DownloadLink    string
	DownloadText    string
}

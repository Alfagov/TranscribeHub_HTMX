package models

import "html/template"

type News struct {
	Id              string        `json:"id,omitempty"`
	SvgNotification template.HTML `json:"svgNotification,omitempty"`
	Title           string        `json:"title,omitempty"`
	Date            string        `json:"date,omitempty"`
	Text            string        `json:"text,omitempty"`
	DownloadLink    string        `json:"downloadLink,omitempty"`
	DownloadText    string        `json:"downloadText,omitempty"`
}

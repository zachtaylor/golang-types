package types

import "html/template"

// NewTemplate is html/template.New
func NewTemplate(string string) *template.Template {
	return template.New(string)
}

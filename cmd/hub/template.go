package main

import (
	"errors"
	"html/template"

	"github.com/jwowillo/hub/cache"
)

// TemplateCache stores parsed template.Templates.
type TemplateCache cache.Cache

// GetTemplate from the TemplateCache or get a new one with Template if
// necessary.
func GetTemplate(c TemplateCache, path string) (*template.Template, error) {
	tmpl := cache.Get(c, path, Template)
	if tmpl == nil {
		return nil, errors.New("couldn't parse template")
	}
	return tmpl.(*template.Template), nil
}

// Template reads the template.Template at the path.
//
// Returns nil if the template.Template doesn't exist or couldn't be parsed.
func Template(path string) interface{} {
	tmpl, err := template.ParseFiles("tmpl/index.html")
	if err != nil {
		return nil
	}
	return tmpl
}

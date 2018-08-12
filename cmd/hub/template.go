package main

import (
	"errors"
	"html/template"

	"gopkg.in/jwowillo/cache.v1"
)

// TemplateCache stores parsed template.Templates.
type TemplateCache cache.Cache

// GetTemplate from the TemplateCache or get a new one with Template if
// necessary.
func GetTemplate(c TemplateCache, path string) (*template.Template, error) {
	tmpl := cache.Get(c, cache.Key(path), TemplateFallback)
	if tmpl == nil {
		return nil, errors.New("couldn't parse template")
	}
	return tmpl.(*template.Template), nil
}

// TemplateFallback adapts Template to a cache.Fallback.
func TemplateFallback(k cache.Key) cache.Value {
	t, err := Template(string(k))
	if err != nil {
		return nil
	}
	return t
}

// Template reads the template.Template at the path.
//
// Returns an error if the template.Template doesn't exist or couldn't be
// parsed.
func Template(path string) (*template.Template, error) {
	return template.ParseFiles("tmpl/index.html")
}

package main

import (
	"html/template"

	"github.com/jwowillo/hub/cache"
)

// ReadTmpl reads the template.Template at the key in the cache.Cache or falls
// back to executing the template.Template if the key doesn't exist.
func ReadTmpl(c cache.Cache, key string) (*template.Template, error) {
	cached, exists := c.Get("tmpl/index.html")
	if !exists {
		tmpl, err := template.ParseFiles("tmpl/index.html")
		if err != nil {
			return nil, err
		}
		c.Put("tmpl/index.html", tmpl)
		cached = tmpl
	}
	return cached.(*template.Template), nil
}

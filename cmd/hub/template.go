package main

import (
	"errors"
	"html/template"

	"github.com/jwowillo/cache/v2"
	"github.com/jwowillo/cache/v2/standard"
)

// TemplateGetter gets a template.Template at a path.
type TemplateGetter func(string) (*template.Template, error)

// MakeTemplateGetter creates the TemplateGetter.
func MakeTemplateGetter() TemplateGetter {
	return MakeTemplateGetterFromGetter(cache.NewFallbackGetter(
		standard.ChangedCache("template"),
		MakeGetterFromTemplateGetter(Template)))
}

// MakeTemplateGetterFromGetter adapts a cache.Getter to a TemplateGetter.
func MakeTemplateGetterFromGetter(g cache.Getter) TemplateGetter {
	return func(p string) (*template.Template, error) {
		x := g.Get(cache.Key(p))
		if x == nil {
			return nil, errors.New("couldn't parse template")
		}
		return x.(*template.Template), nil
	}
}

// MakeGetterFromTemplateGetter adapts a TemplateGetter to a cache.Getter.
func MakeGetterFromTemplateGetter(tg TemplateGetter) cache.Getter {
	return cache.GetterFunc(func(k cache.Key) cache.Value {
		t, err := tg(string(k))
		if err != nil {
			return nil
		}
		return t
	})
}

// Template reads the template.Template at the path.
//
// Returns an error if the template.Template doesn't exist or couldn't be
// parsed.
func Template(path string) (*template.Template, error) {
	return template.ParseFiles("tmpl/index.html")
}

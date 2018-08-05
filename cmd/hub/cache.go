package main

import "html/template"

// TemplateCache stores HTML templates.
type TemplateCache struct {
	data map[string]*template.Template
}

// NewTemplateCache with no stored templates.
func NewTemplateCache() *TemplateCache {
	c := &TemplateCache{}
	c.Clear()
	return c
}

// Get template at the key.
//
// Returns true if the template exists.
func (c TemplateCache) Get(key string) (*template.Template, bool) {
	tmpl, ok := c.data[key]
	return tmpl, ok
}

// Put template at the key.
func (c *TemplateCache) Put(key string, tmpl *template.Template) {
	c.data[key] = tmpl
}

// Clear the cache.
func (c *TemplateCache) Clear() {
	c.data = make(map[string]*template.Template)
}

// FaviconCache stores favicons for Websites.
type FaviconCache struct {
	data map[Website]string
}

// NewFaviconCache with no stored favicons.
func NewFaviconCache() *FaviconCache {
	c := &FaviconCache{}
	c.Clear()
	return c
}

// Get favicon at the key.
//
// Returns true if the favicon exists.
func (c FaviconCache) Get(key Website) (string, bool) {
	f, ok := c.data[key]
	return f, ok
}

// Put favicon at the key.
func (c *FaviconCache) Put(key Website, favicon string) {
	c.data[key] = favicon
}

// Clear the cache.
func (c *FaviconCache) Clear() {
	c.data = make(map[Website]string)
}

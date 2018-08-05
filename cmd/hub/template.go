package main

import (
	"html/template"
	"sync"
)

// TmplCacheMutex is a read/write mutex for updating the template cache.
var TmplCacheMutex = sync.RWMutex{}

// TmplCache to store read templates.
var TmplCache = NewTemplateCache()

// ReadTmpl using TmplCache at key or falling back to executing the template if
// it doesn't exist.
func ReadTmpl(key string) (*template.Template, error) {
	TmplCacheMutex.RLock()
	cached, exists := TmplCache.Get("tmpl/index.html")
	TmplCacheMutex.RUnlock()
	if !exists {
		tmpl, err := template.ParseFiles("tmpl/index.html")
		if err != nil {
			return nil, err
		}
		TmplCacheMutex.Lock()
		TmplCache.Put("tmpl/index.html", tmpl)
		TmplCacheMutex.Unlock()
		cached = tmpl
	}
	return cached, nil
}

// ClearTmplCache clears TmplCache.
func ClearTmplCache() {
	TmplCacheMutex.Lock()
	defer TmplCacheMutex.Unlock()
	TmplCache.Clear()
}

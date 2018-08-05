package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)

// FavCacheMutex is a read/write mutex for updating the favicon cache.
var FavCacheMutex = sync.RWMutex{}

// FavCache to store favicons.
var FavCache = NewFaviconCache()

// Favicon for the Website.
//
// Does nothing if the Website can't be reached or a favicon can't be found.
func Favicon(w Website) string {
	resp, err := http.Get(w.URL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	match := WebsiteRegex.Find(bs)
	if match == nil {
		return ""
	}
	matches := URLRegex.FindAllSubmatch(match, 1)
	if len(matches) == 0 {
		return ""
	}
	match = matches[0][1]
	var favicon string
	if match[0] == '/' {
		favicon = w.URL + string(match)
	} else {
		favicon = string(match)
	}
	return favicon
}

// ReadFavicon using FavCache at key or falling back to making requests to
// get the favicon if it doesn't exist.
func ReadFavicon(w Website) string {
	FavCacheMutex.RLock()
	cached, exists := FavCache.Get(w)
	FavCacheMutex.RUnlock()
	if !exists {
		favicon := Favicon(w)
		FavCacheMutex.Lock()
		FavCache.Put(w, favicon)
		FavCacheMutex.Unlock()
		cached = favicon
	}
	return cached
}

// ClearFavCache clears FavCache.
func ClearFavCache() {
	FavCacheMutex.Lock()
	defer FavCacheMutex.Unlock()
	FavCache.Clear()
}

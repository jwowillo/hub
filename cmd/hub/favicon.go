package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jwowillo/hub/cache"
)

// Favicon for the website at the URL.
//
// Does nothing if the website can't be reached or a favicon can't be found.
func Favicon(u string) string {
	resp, err := http.Get(u)
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
		parsed, err := url.Parse(u)
		if err != nil {
			return ""
		}
		favicon = strings.TrimRight(u, parsed.Path) + string(match)
	} else {
		favicon = string(match)
	}
	return favicon
}

// ReadFavicon reads the favicon at the key in the cache.Cache or falls back to
// making requests to get the favicon if the key doesn't exist.
func ReadFavicon(c cache.Cache, u string) string {
	cached, exists := c.Get(u)
	if !exists {
		favicon := Favicon(u)
		c.Put(u, favicon)
		cached = favicon
	}
	return cached.(string)
}

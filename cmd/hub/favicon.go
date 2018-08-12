package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"gopkg.in/jwowillo/cache.v1"
)

// FaviconCache stores favicons.
type FaviconCache cache.Cache

// GetFavicon from the FaviconCache or get a new one with Favicon if necessary.
func GetFavicon(c FaviconCache, u string) string {
	return cache.Get(c, cache.Key(u), FaviconFallback).(string)
}

// FaviconFallback adapts Favicon to a cache.Fallback.
func FaviconFallback(k cache.Key) cache.Value {
	return Favicon(string(k))
}

// Favicon for the website at the URL.
//
// Returns the empty string if the website can't be reached or a favicon can't
// be found.
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
	match := websiteRegex.Find(bs)
	if match == nil {
		return ""
	}
	matches := urlRegex.FindAllSubmatch(match, 1)
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

// websiteRegex to match favicons with.
var websiteRegex = regexp.MustCompile("<.*rel=\".*icon.*\".*>")

// urlRegex matches URLs.
var urlRegex = regexp.MustCompile("href=\"(.*?)\"")

package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/jwowillo/cache/v2"
	"github.com/jwowillo/cache/v2/standard"
)

// FaviconGetter gets a favicon at a URL.
type FaviconGetter func(string) string

// MakeFaviconGetter creates the FaviconGetter.
func MakeFaviconGetter() FaviconGetter {
	duration := time.Duration(*cacheDuration) * time.Hour
	return MakeFaviconGetterFromGetter(cache.NewFallbackGetter(
		standard.TimeCache("favicon", duration),
		MakeGetterFromFaviconGetter(Favicon)))
}

// MakeFaviconGetterFromGetter adapts a cache.Getter to a FaviconGetter.
func MakeFaviconGetterFromGetter(g cache.Getter) FaviconGetter {
	return func(u string) string {
		return g.Get(cache.Key(u)).(string)
	}
}

// MakeGetterFromFaviconGetter adapts a FaviconGetter to a cache.Getter.
func MakeGetterFromFaviconGetter(fg FaviconGetter) cache.Getter {
	return cache.GetterFunc(func(k cache.Key) cache.Value {
		return fg(string(k))
	})
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

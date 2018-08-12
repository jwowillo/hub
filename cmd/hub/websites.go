package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/jwowillo/cache.v1"

	"gopkg.in/yaml.v2"
)

// WebsitesCache stores parsed Websites.
type WebsitesCache cache.Cache

// GetWebsites from the WebsitesCache or get new ones with Websites if
// necessary.
func GetWebsites(c WebsitesCache, path string) ([]Website, error) {
	ws := cache.Get(c, cache.Key(path), WebsitesFallback)
	if ws == nil {
		return nil, errors.New("couldn't parse websites")
	}
	return ws.([]Website), nil
}

// WebsitesFallback adapts Websites to a cache.Fallback.
func WebsitesFallback(k cache.Key) cache.Value {
	ws, err := Websites(string(k))
	if err != nil {
		return nil
	}
	return ws
}

// Website in directory.
type Website struct {
	URL     string `yaml:"URL"`
	Name    string `yaml:"name"`
	Favicon string
}

// Websites from the config at the path.
func Websites(path string) ([]Website, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var sites []Website
	if err := yaml.Unmarshal(bs, &sites); err != nil {
		return nil, err
	}
	return sites, nil
}

package main

import (
	"errors"
	"io/ioutil"

	"github.com/jwowillo/hub/cache"

	"gopkg.in/yaml.v2"
)

// WebsitesCache stores parsed Websites.
type WebsitesCache cache.Cache

// GetWebsites from the WebsitesCache or get new ones with Websites if
// necessary.
func GetWebsites(c WebsitesCache, path string) ([]Website, error) {
	ws := cache.Get(c, path, Websites)
	if ws == nil {
		return nil, errors.New("couldn't parse websites")
	}
	return ws.([]Website), nil
}

// Website in directory.
type Website struct {
	URL     string `yaml:"URL"`
	Name    string `yaml:"name"`
	Favicon string
}

// Websites ...
func Websites(path string) interface{} {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	var sites []Website
	if err := yaml.Unmarshal(bs, &sites); err != nil {
		return nil
	}
	return sites
}

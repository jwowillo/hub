package main

import (
	"errors"
	"io/ioutil"

	"gopkg.in/jwowillo/cache.v2"
	"gopkg.in/jwowillo/cache.v2/standard"

	"gopkg.in/yaml.v2"
)

// WebsitesGetter gets a list of Websites from a config file.
type WebsitesGetter func(string) ([]Website, error)

// MakeWebsitesGetter creates the WebsitesGetter.
func MakeWebsitesGetter() WebsitesGetter {
	return MakeWebsitesGetterFromGetter(cache.NewFallbackGetter(
		standard.ChangedCache("websites"),
		MakeGetterFromWebsitesGetter(Websites)))
}

// MakeWebsitesGetterFromGetter adapts a cache.Getter to a WebsitesGetter.
func MakeWebsitesGetterFromGetter(g cache.Getter) WebsitesGetter {
	return func(p string) ([]Website, error) {
		x := g.Get(cache.Key(p))
		if x == nil {
			return nil, errors.New("couldn't parse websites")
		}
		return x.([]Website), nil
	}
}

// MakeGetterFromWebsitesGetter adapts a WebsitesGetter to a cache.Getter.
func MakeGetterFromWebsitesGetter(wg WebsitesGetter) cache.Getter {
	return cache.GetterFunc(func(k cache.Key) cache.Value {
		ws, err := wg(string(k))
		if err != nil {
			return nil
		}
		return ws
	})
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

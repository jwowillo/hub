package main

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

// Path to config.
const Path = "config.yaml"

// WebsiteRegex to match favicons with.
var WebsiteRegex = regexp.MustCompile("<.*rel=\".*icon.*\".*>")

// URLRegex matches URLs.
var URLRegex = regexp.MustCompile("href=\"(.*?)\"")

// Website in directory.
type Website struct {
	URL     string `yaml:"URL"`
	Name    string `yaml:"name"`
	Favicon string
}

// ReadConfig at path into Websites.
func ReadConfig(path string) ([]Website, error) {
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

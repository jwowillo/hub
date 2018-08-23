package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// main starts server which serves nothing on the set port.
func main() {
	const configPath = "config.yaml"
	const templatePath = "tmpl/index.html"
	http.Handle("/", MakeHomeHandler(configPath, templatePath))

	const staticPath = "static"
	http.Handle(
		fmt.Sprintf("/%s/", staticPath),
		MakeStaticHandler(staticPath))

	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

// init reads the optional port flag.
func init() {
	flag.Parse()
}

// port to listen on.
var port = flag.Int("port", 8080, "port to listen on")

// cacheDuration before clearing cache.
var cacheDuration = flag.Int(
	"cache-duration",
	24,
	"hours before clearing cache")

package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jwowillo/hub/cache"
)

// FaviconCache ...
type FaviconCache cache.Cache

// ConfigCache ...
type ConfigCache cache.Cache

// TemplateCache ...
type TemplateCache cache.Cache

// Handler returns the main http.HandlerFunc.
func Handler(fc FaviconCache, ConfigCache, tc TemplateCache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := ReadConfig(Path)
		if err != nil {
			log.Println(err)
			return
		}
		var wg sync.WaitGroup
		wg.Add(len(ws))
		for i := range ws {
			x := i
			go func() {
				ws[x].Favicon = ReadFavicon(fc, ws[x].URL)
				wg.Done()
			}()
		}
		wg.Wait()
		tmpl, err := ReadTmpl(tc, "tmpl/index.html")
		if err != nil {
			log.Println(err)
			return
		}
		if err := tmpl.Execute(w, ws); err != nil {
			log.Println(err)
		}
	}
}

// main starts server which serves nothing on the set port.
func main() {
	faviconCache := cache.DefaultTimeCache("favicon",
		time.Duration(*cacheDuration)*time.Hour)
	configCache := cache.DefaultModifiedCache("config")
	templateCache := cache.DefaultModifiedCache("template")
	http.HandleFunc("/", Handler(faviconCache, configCache, templateCache))

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)

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
	"hours before clearing cache",
)

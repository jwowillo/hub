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

// Handler returns the main http.HandlerFunc after injecting all dependencies.
func Handler(fc FaviconCache, wc WebsitesCache, tc TemplateCache,
	configPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := GetWebsites(wc, configPath)
		if err != nil {
			log.Println(err)
			return
		}
		var wg sync.WaitGroup
		wg.Add(len(ws))
		for i := range ws {
			x := i
			go func() {
				ws[x].Favicon = GetFavicon(fc, ws[x].URL)
				wg.Done()
			}()
		}
		wg.Wait()
		tmpl, err := GetTemplate(tc, "tmpl/index.html")
		if err != nil {
			log.Println(err)
			return
		}
		if err := tmpl.Execute(w, ws); err != nil {
			log.Println(err)
		}
	}
}

// Path to config.
const Path = "config.yaml"

// main starts server which serves nothing on the set port.
func main() {
	faviconCache := cache.DefaultTimeCache("favicon",
		time.Duration(*cacheDuration)*time.Hour)
	configCache := cache.DefaultModifiedCache("config")
	templateCache := cache.DefaultModifiedCache("template")
	handler := Handler(faviconCache, configCache, templateCache, Path)

	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	http.HandleFunc("/", handler)
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

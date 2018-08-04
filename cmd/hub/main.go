package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"gopkg.in/yaml.v2"
)

// CacheMutex is a read/write mutex for updating the template cache.
var CacheMutex = sync.RWMutex{}

// Cache to store read templates.
var Cache = NewTmplCache()

// Path to config.
const Path = "config.yaml"

// FaviconRegex to match favicons with.
var FaviconRegex = regexp.MustCompile("<.*rel=\".*icon.*\".*>")

// URLRegex matches URLs.
var URLRegex = regexp.MustCompile("href=\"(.*?)\"")

// Website in directory.
type Website struct {
	URL     string `yaml:"URL"`
	Name    string `yaml:"name"`
	Favicon string
}

// LoadFavicon into the Website.
//
// Does nothing if the Website can't be reached or a favicon can't be found.
func LoadFavicon(w *Website) {
	resp, err := http.Get(w.URL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	match := FaviconRegex.Find(bs)
	if match == nil {
		return
	}
	matches := URLRegex.FindAllSubmatch(match, 1)
	if len(matches) == 0 {
		return
	}
	match = matches[0][1]
	if match[0] == '/' {
		w.Favicon = w.URL + string(match)
	} else {
		w.Favicon = string(match)
	}
}

// ReadConfig at path into Websites.
func ReadConfig(path string) ([]*Website, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var sites []*Website
	if err := yaml.Unmarshal(bs, &sites); err != nil {
		return nil, err
	}
	return sites, nil
}

// Handler serves nothing.
func Handler(w http.ResponseWriter, r *http.Request) {
	ws, err := ReadConfig(Path)
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(ws))
	for _, w := range ws {
		x := w
		go func() {
			LoadFavicon(x)
			wg.Done()
		}()
	}
	wg.Wait()
	tmpl, err := ReadTmpl("tmpl/index.html")
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w, ws); err != nil {
		log.Println(err)
	}
}

// ReadTmpl using Cache at key.
func ReadTmpl(key string) (*template.Template, error) {
	CacheMutex.RLock()
	cached, exists := Cache.Get("tmpl/index.html")
	CacheMutex.RUnlock()
	if !exists {
		tmpl, err := template.ParseFiles("tmpl/index.html")
		if err != nil {
			return nil, err
		}
		CacheMutex.Lock()
		Cache.Put("tmpl/index.html", tmpl)
		CacheMutex.Unlock()
		cached = tmpl
	}
	return cached, nil
}

// ClearCache every set time interval forever.
func ClearCache() {
	go func() {
		for {
			func() {
				CacheMutex.Lock()
				defer CacheMutex.Unlock()
				Cache.Clear()
			}()
			time.Sleep(time.Duration(*cacheDuration) * time.Hour)
		}
	}()
}

// main starts server which serves nothing on the set port.
func main() {
	ClearCache()
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)
	http.HandleFunc("/", Handler)
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

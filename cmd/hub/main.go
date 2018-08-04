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

	"gopkg.in/yaml.v2"
)

// Path to config.
const Path = "config.yaml"

// Tmpl to inject website directory into.
var Tmpl = template.Must(template.New("directory").Parse(`<!doctype html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <title>hub</title>

  <link rel="stylesheet" href="/static/styles.css" />
</head>

<body>
  <header>
    <h1>hub</h1>
  </header>

  <ul>
    {{ range . }}
    <li>
      <img src="{{ .Favicon }}" />
      <a href="{{ .URL }}" target="_blank">
        <span>{{ .Name }}</span>
      </a>
    </li>
    {{ end }}
  </ul>
</body>

</html>`))

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
	if err := Tmpl.Execute(w, ws); err != nil {
		log.Println(err)
	}
}

// main starts server which serves nothing on the set port.
func main() {
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

package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Path to config.
const Path = "config.yaml"

// Tmpl to inject website directory into.
const Tmpl = `<!doctype html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <title>hub</title>
</head>

<body>
  <ul>
    {{ range . }}
    <li>
      <a href="{{ .URL }}" target="_blank">{{ .Name }}</a>
    </li>
    {{ end }}
  </ul>
</body>

</html>`

// Website in directory.
type Website struct {
	URL  string `yaml:"URL"`
	Name string `yaml:"name"`
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

// Handler serves nothing.
func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("directory").Parse(Tmpl)
	if err != nil {
		log.Println(err)
		return
	}
	ws, err := ReadConfig(Path)
	if err != nil {
		log.Println(err)
		return
	}
	if err := tmpl.Execute(w, ws); err != nil {
		log.Println(err)
	}
}

// main starts server which serves nothing on the set port.
func main() {
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

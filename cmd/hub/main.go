package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

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

// Handler serves nothing.
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
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

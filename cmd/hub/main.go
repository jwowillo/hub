package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler serves nothing.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}

// main starts server which serves nothing on the set port.
func main() {
	http.HandleFunc("/", handler)
	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}

// init reads the optional port flag.
func init() {
	flag.Parse()
}

// port to listen on.
var port = flag.Int("port", 8080, "port to listen on")

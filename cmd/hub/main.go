package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Handler serves nothing.
func Handler(w http.ResponseWriter, r *http.Request) {
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
			ws[x].Favicon = ReadFavicon(ws[x])
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

// main starts server which serves nothing on the set port.
func main() {
	Schedule(ClearTmplCache, time.Duration(*cacheDuration))
	Schedule(ClearFavCache, time.Duration(*cacheDuration))
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

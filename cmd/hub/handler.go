package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// MakeHomeHandler makes the home page http.Handler.
func MakeHomeHandler(configPath, templatePath string) http.Handler {
	return http.HandlerFunc(Handler(
		MakeFaviconGetter(), MakeWebsitesGetter(), MakeTemplateGetter(),
		configPath, templatePath))
}

// MakeStaticHandler makes the static-file http.Handler.
func MakeStaticHandler(staticPath string) http.Handler {
	return http.StripPrefix(
		fmt.Sprintf("/%s/", staticPath),
		http.FileServer(http.Dir(staticPath)))
}

// Handler returns the main http.HandlerFunc after injecting all dependencies.
func Handler(
	faviconGetter FaviconGetter,
	websitesGetter WebsitesGetter,
	templateGetter TemplateGetter,
	configPath, templatePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := websitesGetter(configPath)
		if err != nil {
			log.Println(err)
			return
		}
		var wg sync.WaitGroup
		wg.Add(len(ws))
		for i := range ws {
			x := i
			go func() {
				ws[x].Favicon = faviconGetter(ws[x].URL)
				wg.Done()
			}()
		}
		wg.Wait()
		tmpl, err := templateGetter(templatePath)
		if err != nil {
			log.Println(err)
			return
		}
		if err := tmpl.Execute(w, ws); err != nil {
			log.Println(err)
		}
	}
}

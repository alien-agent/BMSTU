package main

import (
	"github.com/mmcdole/gofeed"
	"html/template"
	"log"
	"net/http"
	"regexp"
)

var cdataRemoverRegexp = regexp.MustCompile("<img.*>")
var sitesMap = map[string]string{
	"blagnews": "http://blagnews.ru/rss_vk.xml",
	"lenta":    "https://lenta.ru/rss",
	"vz":       "https://vz.ru/rss.xml",
}

func HomeRouterHandler(w http.ResponseWriter, req *http.Request) {
	link, ok := sitesMap[req.RequestURI[1:]]
	if !ok {
		link = "http://blagnews.ru/rss_vk.xml"
	}
	log.Println("Handling request from", req.RemoteAddr)

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(link)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Remove CDATA from description as we render image provided from Item.Enclosures[0].URL
	for _, item := range feed.Items {
		item.Description = cdataRemoverRegexp.ReplaceAllString(item.Description, "")
	}

	templates, _ := template.ParseFiles("index.html")
	t := templates.Lookup("index.html")
	t.Execute(w, feed)

	log.Println("Template rendered successfully")
}

func main() {
	http.HandleFunc("/", HomeRouterHandler)
	log.Println("Listening http://localhost:9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

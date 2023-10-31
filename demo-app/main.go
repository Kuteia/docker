package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	accentColor := os.Getenv("ACCENT_COLOR")
	if accentColor == "" {
		accentColor = "#f39200"
	}
	hostname, _ := os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	http.HandleFunc("/", defaultHandler(accentColor, hostname))

	err := http.ListenAndServe(":80", nil)
	log.Fatal(err)
}

func defaultHandler(accentColor, hostname string) http.HandlerFunc {
	t := template.Must(template.ParseFiles("index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, struct{ AccentColor, Hostname string }{accentColor, hostname})
	}
}
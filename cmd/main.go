package main

import (
	"log"
	"net/http"
	"os"

	"presthus.dev/views/pages"

	"github.com/a-h/templ"
)

func main() {
    port := os.Getenv("PORT")

    if port == "" {
        log.Fatal("$PORT must be set")
    }

    http.Handle("/", templ.Handler(pages.Index()))

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Fatal(http.ListenAndServe(":" + port, nil))
}


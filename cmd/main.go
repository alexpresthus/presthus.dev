package main

import (
	"fmt"
	"log"
	"net/http"

	"presthus.dev/views/pages"

	"github.com/a-h/templ"
)

func main() {
    fmt.Println("Server up")

    http.Handle("/", templ.Handler(pages.Index()))

    fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Fatal(http.ListenAndServe(":8080", nil))
}


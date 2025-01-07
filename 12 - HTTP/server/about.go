package main

import (
	"fmt"
	"net/http"
	"os"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	content, _ := os.ReadFile("about.html")
	fmt.Fprintf(w, "%s", string(content))
}

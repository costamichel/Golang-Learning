package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Go HTTP!")
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("nome")

	if name == "" {
		name = "Strange"
	}

	fmt.Fprintf(w, "Hello %s, welcome to Michel's server!", name)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("15:04:05")
	fmt.Fprintf(w, "Horário de agora: %s", formattedTime)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	})
	http.HandleFunc("/saudacao", greetingHandler)
	http.HandleFunc("/sobre", AboutHandler)
	http.HandleFunc("/hora", timeHandler)

	fmt.Println("Starting Server localhost at por 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Couldn't start server!")
		return
	}
}

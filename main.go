package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"sync"
)

func main() {
	fmt.Println("Hello, world")
	var wg sync.WaitGroup
	go noAbstraction(&wg)
	wg.Wait()
}

func noAbstraction(wg *sync.WaitGroup) {
	wg.Add(1)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	wg.Done()
}

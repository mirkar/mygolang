package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8082
	apphandler := http.FileServer(http.Dir("./static"))
	http.Handle("/", apphandler)
	http.Handle("/app/", http.StripPrefix("/app/", apphandler))
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
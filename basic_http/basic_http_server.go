package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloJsonHandler struct{}

var handler http.Handler

func main() {
	port := 8082

	http.HandleFunc("/helloworld", helloWorldHandler)

	handler = newHelloJsonHandler()
	http.Handle("/hellojson", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World\n")
	handler.ServeHTTP(w, r)
}

func newHelloJsonHandler() http.Handler {
	return helloJsonHandler{}
}

func (h helloJsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Jason\n")
}

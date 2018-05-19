// Command staticsite is a trivial web server that serves static content
// placed under directory with name static. 
// 
// It should be executed from the directory in which the source resides,
// as it will look for "static" directory  and its files in the current directory.
//
// Content access
//
// For testing purposes an access to the static content is possible in two different ways:
// simply from the root / and as /app/.
//     apphandler := http.FileServer(http.Dir("./static"))
//     http.Handle("/app/", http.StripPrefix("/app/", apphandler))
//
//https://stackoverflow.com/questions/27945310/why-do-i-need-to-use-http-stripprefix-to-access-my-static-files
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
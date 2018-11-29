// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// testserver is a trivial web server mostly to play with docker and kibana
// any request to the server returns plain text listing all http headers from the request.
//
//
// Usage:
//
//	testserver  [--port -port] <num>
//	testserver  [--port -port]=<num>
//
// Example:
// 	> nohup testserver  -port 8082 > log.txt&
//	> curl http://localhos:8082/
// 	GET / HTTP/1.1
//	Header field "Cookie", Value ["Idea-4290136c=ce274889-7ffb-4f58-a5f8-95a8c73d4aaa"]
//	Header field "Connection", Value ["keep-alive"]
//	Header field "Upgrade-Insecure-Requests", Value ["1"]
//	Header field "User-Agent", Value ["Mozilla/5.0 (Windows NT 6.1; WOW64; rv:60.0) Gecko/20100101 Firefox/60.0"]
//	Header field "Accept", Value ["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"]
//	Header field "Accept-Language", Value ["en-US,en;q=0.5"]
//	Header field "Accept-Encoding", Value ["gzip, deflate"]
//	Host = "localhost:8082"
//	RemoteAddr= "127.0.0.1:61510"
//	Finding value of "Accept" ["text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"]
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// BUG(r): The rule Title uses for word boundaries does not handle Unicode punctuation properly.
func main() {
	// file system roots
	// TODO(gri) consider the invariant that goroot always end in '/'
	port := flag.String("port", ":8082", "port number")
	flag.Parse()
	fmt.Println("port", *port)

	http.HandleFunc("/", reqDetailsHandler)

	log.Printf("Server starting on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s", *port), nil))
}

func reqDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
	//Iterate over all header fields
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)
	//Get value for a specified token
	fmt.Fprintf(w, "\n\nFinding value of \"Accept\" %q", r.Header["Accept"])

}

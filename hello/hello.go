// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// godoc2md converts godoc formatted package documentation into Markdown format.
//
//
// Usage
//
//    godoc2md $PACKAGE > $GOPATH/src/$PACKAGE/README.md
//
// Another line with description and maybe additional usage sample
//    hello $SOMEPARAM
package main

import (
	"fmt"
	"github.com/mirkar/mygolang/hello/foobar"
)

func main() {
	fmt.Printf(foobar.Reverse("\n!oG ,olleH"))
	fmt.Printf(foobar.Reverse("\nhello, world"))
	fmt.Printf("hello, world\n")
}

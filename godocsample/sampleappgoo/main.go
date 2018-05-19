/*
Copyright 2016 Google Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
     http://www.apache.org/licenses/LICENSE-2.0
     Unless required by applicable law or agreed to in writing, software
     distributed under the License is distributed on an "AS IS" BASIS,
     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
     See the License for the specific language governing permissions and
     limitations under the License.
*/

// Command template is a trivial web server that uses the text/template (and
// html/template) package's "block" feature to implement a kind of template
// inheritance.
//
// It should be executed from the directory in which the source resides,
// as it will look for its template files in the current directory.
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, sample app goo\n")
}

// Copyright 2017 Frank Schroeder. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package properties provides functions for reading and writing
// ISO-8859-1 and UTF-8 encoded .properties files and has
// support for recursive property expansion.
//
// Java properties files are ISO-8859-1 encoded and use Unicode
// literals for characters outside the ISO character set. Unicode
// literals can be used in UTF-8 encoded properties files but
// aren't necessary.
//
// To load a single properties file use MustLoadFile():
//
//   p := properties.MustLoadFile(filename, properties.UTF8)
//
package godocsample

// jsonpretty is a small tool to prettify json-documents
//
// author: mathias gumz <mg@2hoch5.com>
// date:   2013-04-02
// license:
//
//   Copyright (c) 2013, Mathias Gumz
//   All rights reserved.
//
//   Redistribution and use in source and binary forms, with or without
//   modification, are permitted provided that the following conditions are met:
//
//   Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//   Redistributions in binary form must reproduce the above copyright notice, this
//   list of conditions and the following disclaimer in the documentation and/or
//   other materials provided with the distribution.
//   THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
//   AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
//   IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
//   DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
//   FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
//   DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
//           SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
//   HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
//   LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT
//   OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH
//   DAMAGE.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func onErrorExit(e error, exit int) {
	if e != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %s\n", e))
		os.Exit(exit)
	}
}

func main() {

	var (
		in     = flag.String("in", "", "name of input.json, if empty <stdin> is used")
		out    = flag.String("out", "", "name of output.json, if empty <stdout> is used")
		indent = flag.String("indent", "  ", "string to use as indention")
		prefix = flag.String("prefix", "", "string to use as prefix")
		err    error
	)

	flag.Parse()

	infile := os.Stdin
	if len(*in) > 0 {
		infile, err = os.Open(*in)
		onErrorExit(err, 1)
	}

	outfile := os.Stdout
	if len(*out) > 0 {
		outfile, err = os.Create(*out)
		onErrorExit(err, 1)
	}

	json_in, err := ioutil.ReadAll(infile)
	onErrorExit(err, 1)

	buf := bytes.Buffer{}
	err = json.Indent(&buf, json_in, *prefix, *indent)
	onErrorExit(err, 1)

	outfile.Write(buf.Bytes())
	outfile.WriteString("\n")
}

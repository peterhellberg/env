# env

[![Build Status](https://travis-ci.org/peterhellberg/env.svg?branch=master)](https://travis-ci.org/peterhellberg/env)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/env)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/env#license-mit)

Load environment variables into Go types, with fallback values.

This package is meant to be used when configuring a [twelve-factor app](http://12factor.net/).

Currently supported types are `bool`, `[]byte`, `time.Duration`, `float64`, `int`, `string`, `[]string` and `*url.URL`

## Installation

    go get -u github.com/peterhellberg/env

Feel free to copy all or parts of this package into your own codebase.

## Usage

```go
package main

import (
	"fmt"

	"github.com/peterhellberg/env"
)

func main() {
	fmt.Println(
		env.Bool("BOOL", false),
		env.Bytes("BYTES", []byte{4, 2}),
		env.Duration("DURATION", 250000),
		env.Float64("FLOAT64", float64(2.5)),
		env.Int("INT", 1337),
		env.String("STRING", "Foobar"),
		env.Strings("STRINGS", []string{"Foo", "Bar"}),
		env.URL("URL", &url.URL{Scheme: "http", Host: "example.com"}),
	)
}
```

```bash
$ go run example.go
false [4 2] 250µs 2.5 1337 Foobar [Foo Bar] http://example.com

$ BOOL=true BYTES=foo DURATION=24m FLOAT64=11.2 INT=2600 STRING=hello STRINGS=a,b URL=http://c7.se/ go run example.go
true [102 111 111] 24m0s 11.2 2600 hello [a b] http://c7.se/
```

## License (MIT)

Copyright (c) 2015-2016 [Peter Hellberg](http://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:

> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

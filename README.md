# Go-WazirX
[![itsksaurabh](https://circleci.com/gh/itsksaurabh/go-wazirx.svg?style=shield)](https://circleci.com/gh/itsksaurabh/workflows/go-wazirx/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsksaurabh/go-wazirx)](https://goreportcard.com/report/github.com/itsksaurabh/go-wazirx)
[![MIT License](https://img.shields.io/github/license/itsksaurabh/go-wazirx?style=social)](https://github.com/itsksaurabh/go-wazirx/blob/master/LICENSE)
___

<img style="float:left;" width="200" src="./assets/logo.png"> 

## Go-WazirX is a [Go](http://golang.org/) client library for accessing the [WazirX](https://wazirx.com/)'s Public Rest API.

## API Documentation
You can read the API server documentation [here](https://github.com/WazirX/wazirx-api).

## Usage
Package provides a client for accessing different endpoints of the API.
Create a new instance of Client, then use the various methods on the client to access different parts of the API.

For demonstration:
```
package main

import (
	"context"
	"fmt"
	"log"
	
	"github.com/itsksaurabh/go-wazirx"
)

func main() {
        // client for accessing different endpoints of the API
	client := wazirx.Client{}
	ctx := context.Background()

	data, err := client.MarketStatus(ctx)
	if err != nil {
		log.Fatal("request failed:", err)
	}
	fmt.Println(data)
}

  ```
Notes:
* Using the  [https://godoc.org/context](https://godoc.org/context) package for passing context.
* Look at tests(*_test.go) files for more sample usage.

# Contributing
I welcome pull requests, bug fixes and issue reports. Before proposing a change, please discuss your change by raising an issue.

# Author
<ul>
  <li><a href="https://in.linkedin.com/in/itsksaurabh">Kumar Saurabh</a></li>
</ul>

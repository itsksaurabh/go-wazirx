# Go-WazirX
[![itsksaurabh](https://circleci.com/gh/itsksaurabh/go-wazirx.svg?style=shield)](https://circleci.com/gh/itsksaurabh/workflows/go-wazirx/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/itsksaurabh/go-wazirx)](https://goreportcard.com/report/github.com/itsksaurabh/go-wazirx)
[![GoDoc](https://godoc.org/github.com/itsksaurabh/go-wazirx?status.svg)](https://godoc.org/github.com/itsksaurabh/go-wazirx)
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)
[![MIT License](https://img.shields.io/github/license/itsksaurabh/go-wazirx?style=social)](https://github.com/itsksaurabh/go-wazirx/blob/master/LICENSE)
___

<img style="float:left;" width="200" src="./assets/logo.png"> 

## Go-WazirX is a [Go](http://golang.org/) client library for accessing the [WazirX](https://wazirx.com/)'s Public Rest API.
This project has been recommended by the **founders** of [WazirX](https://wazirx.com/) cryptocurrency exchange themselves. You can check [here](https://twitter.com/BuddhaSource/status/1239103029112012800).
## API Documentation
You can read the API server documentation [here](https://github.com/WazirX/wazirx-api).

## Installation

Make sure you have set the environment variable $GOPATH

```bash
export GOPATH="path/to/your/go/folder"
```

Obtain the latest version of the  Go-Wazirx library with:

```bash
go get github.com/itsksaurabh/go-wazirx
```

Then, add the following to your Go project:

```go
import (
	"github.com/itsksaurabh/go-wazirx"
)
```

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

## Error Handling

All errors generated at runtime will be returned to the calling client method. Any API request for which WazirX returns an error encoded in a JSON response will be parsed and returned by the client method as a Golang error struct. Lastly, it is important to note that for HTTP requests, if the response code returned is not '200 OK', an error will be returned to the client method detailing the response code that was received.

## Testing

In order to run the tests for this library, you will first need to install [Mage](https://magefile.org/) - A Make/rake-like dev tool using Go. You can install the dependency with the following command:

**Using GOPATH**

```
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

**Using Go Modules**

```
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```
The mage binary will be created in your `$GOPATH/bin` directory.
You may also install a binary release from Mage's [releases](https://github.com/magefile/mage/releases) page.

Then run all tests by executing the following in your command line:
    
 	$ mage -v Test

**Updating Test Data**

You can update the test data inside `./testdata/` directory by enabling the `-update` flag while testing. By default the flag is set to `false`.

Or

simply Run the following command to update all test data :

```sh
$ mage generate
```

# Contributing
I welcome pull requests, bug fixes and issue reports. Before proposing a change, please discuss your change by raising an issue.

# Maintainer 😎

[Kumar Saurabh](https://in.linkedin.com/in/itsksaurabh)

## License

[MIT](LICENSE) © Kumar Saurabh

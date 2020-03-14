/*
Package wazirx provides a client for using the Wazirx's Public Rest API.
You can read the API server documentation at https://github.com/WazirX/wazirx-api

Usage:

create a new instance of Client, then use the various methods on the client to access different parts of the API.
For demonstration:
  package main
  import (
	"context"
	"fmt"
	"log"

   	"github.com/itsksaurabh/go-wazirx"
 )

  func main() {
	client := wazirx.Client{}
	ctx := context.Background()

	data, err := client.MarketStatus(ctx)
	if err != nil {
		log.Fatal("request failed:", err)
	}
	fmt.Println(data)
  }

Notes:
* Using the  [https://godoc.org/context](https://godoc.org/context) package for passing context.
* Look at tests(*_test.go) files for more sample usage.

*/
package wazirx

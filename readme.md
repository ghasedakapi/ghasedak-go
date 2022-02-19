# ghasedak-go

[Ghasedak sms gateway](https://ghasedak.io) for golang.

## install

```sh
go get github.com/ghasedakapi/ghasedak-go
```

## example

```go
package main

import (
	"fmt"

	ghasedak "github.com/ghasedakapi/ghasedak-go"
)

func main() {

	// initialize connection:
	c := ghasedak.NewClient("api_key", "")

	// Send a new text massage:
	r := c.Send("Hello world!", "09xxxxxxxx")
	fmt.Println(r.Code)
	fmt.Println(r.Message)

	// Send group massages:
	r := c.Bulk1("Hello world!", "09xxxxxxxx,09xxxxxxxx,09xxxxxxxx")
	fmt.Println(r.Code)
	fmt.Println(r.Message)
	// -----------
	r := c.Bulk2("Hello world!", "09xxxxxxxx,09xxxxxxxx,09xxxxxxxx")
	fmt.Println(r.Code)
	fmt.Println(r.Message)

	// Check the status of massages:
	r := c.Status("Massage_ID", "1")
	fmt.Println(r.Message)
	fmt.Println(r.Code)

	// Send verification massages:
	r := c.SendOTP("09xxxxxxxxx", "Your Template", Param1)
	fmt.Println(r.Message)
	fmt.Println(r.Code)
}

```

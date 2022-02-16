ghasedak-go
===============
[Ghasedak sms gateway](https://ghasedak.io) for golang.

install
-------

```sh
go get github.com/ghasedakapi/ghasedak-go
```

example
-------
```go
package main

import (
	"fmt"

	ghasedak "github.com/ghasedakapi/ghasedak-go"
)

func main() {
	c := ghasedak.NewClient("api_key", "xxxxxxxx")

	r := c.Send("Hello, Brave new world!", "09xxxxxxxx")
	fmt.Println(r.Code)
	fmt.Println(r.Message)
}

```
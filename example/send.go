package main

import (
	"fmt"

	ghasedak "github.com/ghasedakapi/ghasedak-go"
)

func main() {
	c := ghasedak.NewClient("api_key", "")

	r := c.Send("Hello, Brave new world!", "0935000000")
	fmt.Println(r.Code)
	fmt.Println(r.Message)
}

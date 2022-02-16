package main

import (
	"fmt"

	ghasedak "github.com/ghasedakapi/ghasedak-go"
)

func main() {
	c := ghasedak.NewClient("05624f7710742db14636cc4700ec5a4464a4a467efb998f80e4820282ec4d5bc", "10008566")

	// sms := c.Send("!", "09120581875")
	// fmt.Println(sms.Code)
	// fmt.Println(sms.Message)
	statuscheck := c.Status("2914845496", "1")
	fmt.Println(statuscheck.Code)
	fmt.Println(statuscheck.Message)
}

package ghasedak

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

// Client type
type Client struct {
	APIKEY     string
	LineNumber string
}

// Response type
type Response struct {
	Success bool
	Code    int
	Message string
	ID      int64
}

// NewClient create new ghasedak client
func NewClient(apikey, linenumber string) Client {
	return Client{APIKEY: apikey, LineNumber: linenumber}
}

// Send simple sms
func (c *Client) Send(msg, receptor string) Response {
	route := "http://api.ghasedak.io/v2/sms/send/simple"
	data := strings.NewReader(fmt.Sprintf(`message=%s&receptor=%s&linenumber=%s`,
		msg, receptor, c.LineNumber))

	rq, err := http.NewRequest("POST", route, data)
	if err != nil {
		log.Println(err)
		return Response{Success: false, Message: err.Error()}
	}
	rq.Header.Set("Cache-Control", "no-cache")
	rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rq.Header.Set("Apikey", c.APIKEY)

	rp, err := http.DefaultClient.Do(rq)
	if err != nil {
		log.Println(err)
		return Response{Success: false, Message: err.Error()}
	}
	defer rp.Body.Close()

	response := Response{}

	if rp.StatusCode == http.StatusOK {
		response.Code = http.StatusOK
		bodyBytes, err := ioutil.ReadAll(rp.Body)
		if err != nil {
			log.Println(err)
			return Response{Success: true, Message: err.Error()}
		}
		bodyString := string(bodyBytes)
		response.Message = gjson.Get(bodyString, "result.message").String()
		response.ID = gjson.Get(bodyString, "items.۰").Int()
		response.Success = true

		return response
	}

	response.Code = rp.StatusCode
	response.Success = false

	return response
}

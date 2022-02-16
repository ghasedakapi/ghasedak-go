package ghasedak

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// Client type
type Client struct {
	APIKEY     string
	LineNumber string
	host       string
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
	return Client{APIKEY: apikey, LineNumber: linenumber, host: "api.ghasedak.me"}
}

func (c *Client) SetHost(host string) {
	c.host = host
}

// get status
func (c *Client) Status(id string, itype string) Response {
	route := "http://" + c.host + "/v2/sms/status?agent=go"
	data := strings.NewReader(fmt.Sprintf(`id=%s&type=%s`,
		id, itype))

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

// Send simple sms
func (c *Client) Send(msg, receptor string) Response {
	route := "http://" + c.host + "/v2/sms/send/simple?agent=go"
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

func (c *Client) SendOTP(receptor string, template string, code int) Response {
	route := "http://" + c.host + "/v2/verification/send/simple?agent=go"

	data := url.Values{}
	data.Set("template", template)
	data.Set("type", "1")
	data.Set("receptor", receptor)
	data.Set("linenumber", c.LineNumber)
	data.Set("param1", strconv.Itoa(code))

	req, err := http.NewRequest("POST", route, bytes.NewBuffer([]byte(data.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Apikey", c.APIKEY)
	if err != nil {
		log.Println(err)
		return Response{Success: false, Message: err.Error()}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("erro is -> %s", err.Error())
		return Response{Success: false, Message: err.Error()}
	}
	defer res.Body.Close()

	response := Response{}

	if res.StatusCode == http.StatusOK {
		response.Code = http.StatusOK
		bodyBytes, err := ioutil.ReadAll(res.Body)
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

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{Success: true, Message: err.Error()}
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	response.Code = res.StatusCode
	response.Message = gjson.Get(bodyString, "result.message").String()
	response.ID = gjson.Get(bodyString, "items.۰").Int()
	response.Success = false

	return response
}

func (c *Client) SendVoice(message string, receptor string, template string) Response {

	route := "http://api.ghasedak.me/v2/verification/send/simple?agent=go"

	v := url.Values{}
	v.Set("template", template)
	v.Set("apikey", c.APIKEY)
	v.Set("receptor", receptor)
	v.Set("message", message)

	req, err := http.NewRequest("POST", route, bytes.NewBuffer([]byte(v.Encode())))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	req.Header.Set("Apikey", c.APIKEY)
	if err != nil {
		log.Println(err)
		return Response{Success: false, Message: err.Error()}
	}
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()

	response := Response{}

	if res.StatusCode == http.StatusOK {
		response.Code = http.StatusOK
		bodyBytes, err := ioutil.ReadAll(res.Body)
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

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{Success: true, Message: err.Error()}
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
	response.Code = res.StatusCode
	response.Message = gjson.Get(bodyString, "result.message").String()
	response.ID = gjson.Get(bodyString, "items.۰").Int()
	response.Success = false

	return response
}

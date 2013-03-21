package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"net/url"
	"bytes"
)

var yHost, fileName, command string

type RequestBody struct {
	*bytes.Buffer
}

type Header struct {
	Heads map[string]string
	content_length	int64
}

func init () {
	flag.StringVar(&yHost, "h", yHost, "yamaha:port")
	flag.StringVar(&command, "c", command, "up or down")
}

func (rb *RequestBody) Close() error {return nil}

func pull(t url.URL, r chan *http.Response) {
	response, err := http.Get(t.String())
	if err != nil {
		log.Fatal(err)
	}
	r <- response
}

func push(t url.URL, body *RequestBody, header Header, r chan *http.Response) {
	req, err := http.NewRequest("POST", t.String(), body)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range(header.Heads){
		req.Header.Set(k, v)
	}
	req.ContentLength = header.content_length

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	r <- resp
}

func makeHeader(cl int64) Header {
	w := map[string]string{
		"User-Agent": "AV_Controller/1.0 (iOS)",
		"SOAPACTION": "\"urn:schemas-upnp-org:service:AVTransport:1#GetTransportInfo\"",
		"Content-Type": "text/xml; charset=\"utf-8\"",
		"Accept-Language": "en-us",
		"Accept-Encoding": "gzip, deflate",
		"Connection": "keep-alive",
		"Proxy-Connection": "keep-alive",
	}
	h := Header{Heads: w, content_length: cl}
	return h
}

func main () {
	flag.Parse()
	target := ""
	rChannel := make(chan *http.Response)
	actionPath := "/YamahaRemoteControl/ctrl"
	buf := bytes.NewBuffer(make([]byte, 0))
	b := make([]byte, 0)
		
	if yHost != "" {
		target = "http://" + yHost + actionPath
	} else {
		//for testing with a proxy
		target = "http://127.0.0.1:8080" + actionPath
	}
	u, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}
	
	switch command {
	case "up":
		b = volumeUp()
	case "down":
		b = volumeDown()
	case "mute":
		b = toggleMute()
	case "on":
		b = powerOn()
	case "off":
		b = standBy()
	default:
		// oh noes. this could be bad.
		b = volumeAt(command)
	}	
	
	buf.Write(b)	
	header := makeHeader(int64(len(b)))
	body := &RequestBody{buf}
	
	go push(*u, body, header, rChannel)
	response := <- rChannel
	fmt.Println(response.Status)
}
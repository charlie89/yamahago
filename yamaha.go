package main

/*
A small command line utility for controlling the Yamaha RX-A720 in my living room.
*/

import (
  "bytes"
  "flag"
  "fmt"
  "log"
  "net/http"
  "net/url"
)

var yHost, fileName, command string

type RequestBody struct {
  *bytes.Buffer
}

type Header struct {
  Heads          map[string]string
  content_length int64
}

func init() {
  flag.StringVar(&yHost, "h", yHost, "yamaha:port")
  flag.StringVar(&command, "c", command, "up, down, mute, on, off or a 3-digit number to set decibels")
}

func (rb *RequestBody) Close() error { return nil }

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
  for k, v := range header.Heads {
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
    "User-Agent":       "AV_Controller/1.0 (iOS)",
    "SOAPACTION":       "\"urn:schemas-upnp-org:service:AVTransport:1#GetTransportInfo\"",
    "Content-Type":     "text/xml; charset=\"utf-8\"",
    "Accept-Language":  "en-us",
    "Accept-Encoding":  "gzip, deflate",
    "Connection":       "keep-alive",
    "Proxy-Connection": "keep-alive",
  }
  h := Header{Heads: w, content_length: cl}
  return h
}

func volumeUp() []byte {
	// 1 decibel is the jnd for the normal human ear
  return []byte("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Volume><Lvl><Val>Up 2 dB</Val><Exp></Exp><Unit></Unit></Lvl></Volume></Main_Zone></YAMAHA_AV>\r\n")
}

func volumeDown() []byte {
  return []byte("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Volume><Lvl><Val>Down 2 dB</Val><Exp></Exp><Unit></Unit></Lvl></Volume></Main_Zone></YAMAHA_AV>\r\n")
}

func volumeAt(str string) []byte {
  m := "<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Volume><Lvl><Val>-%s</Val><Exp>1</Exp><Unit>dB</Unit></Lvl></Volume></Main_Zone></YAMAHA_AV>\r\n"
  return []byte(fmt.Sprintf(m, str))
}

func toggleMute() []byte {
  return []byte("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Volume><Mute>On/Off</Mute></Volume></Main_Zone></YAMAHA_AV>\r\n")
}

func powerOn() []byte {
  return []byte("<YAMAHA_AV cmd=\"PUT\"><System><Power_Control><Power>On</Power></Power_Control></System></YAMAHA_AV>\r\n")
}

func standBy() []byte {
  return []byte("<YAMAHA_AV cmd=\"PUT\"><System><Power_Control><Power>Standby</Power></Power_Control></System></YAMAHA_AV>\r\n")
}

func switchInput(input string) []byte {
  return []byte("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Input><Input_Sel>" + input + "</Input_Sel></Input></Main_Zone></YAMAHA_AV>\r\n")
}


func main() {
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
    target = "http://rx-v679.lan" + actionPath
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
  case "hdmi1":
    b = switchInput("HDMI1")
  case "hdmi2":
    b = switchInput("HDMI2")
  case "hdmi3":
    b = switchInput("HDMI3")
  case "hdmi4":
    b = switchInput("HDMI4")
  case "hdmi5":
    b = switchInput("HDMI5")
  case "netradio":
    b = switchInput("NET RADIO")
  case "server":
    b = switchInput("SERVER")
  default:
    // oh noes. this could be bad.
    // therefore commented it out, not that i blow the neighbors away :)
    //b = volumeAt(command)
  }

  buf.Write(b)
  header := makeHeader(int64(len(b)))
  body := &RequestBody{buf}

  go push(*u, body, header, rChannel)
  response := <-rChannel
  fmt.Println(response.Status)
}

package main

import "fmt"

// start off with messages as strings

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

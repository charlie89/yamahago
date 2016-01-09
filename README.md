yamahago
========

Controls your Yamaha AV/Receiver, should work with most RX-V* or RX-A* AV/Receivers.

The compiled program has very little latency compared to something bigger like yamaha-nodejs.
I use this script to control my AV-Receiver right off the G-Keys on my gaming keyboard.


Turn the volume up 2db:

    yamaha -h 192.168.1.21 -c up

Turn the volume down 2db:

    yamaha -h 192.168.1.21 -c down

Toggle mute/unmute:

    yamaha -h 192.168.1.21 -c mute

Power on: 

    yamaha -h 192.168.1.21 -c on

Power standby:

    yamaha -h 192.168.1.21 -c off

Switch to input HDMI1: (hdmi1-hdmi5 implemented)

    yamaha -h 192.168.1.21 -c hdmi1

Switch to input SERVER:

    yamaha -h 192.168.1.21 -c server

Switch to input NET RADIO:

    yamaha -h 192.168.1.21 -c netradio



Special thanks to boboTjones for this beautiful template, it saved me a lot of work.

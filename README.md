yamahago
========

In the process of writing an HTTPproxy in Go, I started using it to intercept traffic between the iPhone app for the Yamaha RX-A720 that sits in my living room. I used the intercepted requests to build this little controller (and also discover that the thing has a web UI). 

Turn the volume up 2db:

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c up
    200 OK

Turn the volume down 2db:

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c down
    200 OK

Set the volume to -34.0db:

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c 340
    200 OK

Toggle mute/unmute:

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c mute
    200 OK

Power on: 

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c on
    200 OK

Power standby:

    lollipop:go erin$ bin/yamaha -h 192.168.1.83:80 -c off
    200 OK


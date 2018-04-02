package main

import (
  "log"
  "net"
  "os"
  
  "github.com/coreos/go-systemd/activation"
)

func Listeners() (net.Listener, net.Listener) {
  listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}
	
	switch len(listeners) {
	  case 2:
      return listeners[0], listeners[1]
	  case 0:
	    log.Print("Systemd not activated, creating sockets")
	    unix, err := net.Listen("unixpacket","/tmp/tasks.sock")
    	if err != nil {
    		panic(err)
    	}
    	defer os.Remove("/tmp/tasks.sock")
    	
	    tcp, err := net.Listen("tcp","9000")
    	if err != nil {
    		panic(err)
    	}
	    return unix, tcp
    default:
		  panic("Unexpected number of socket activation fds")
	}
}
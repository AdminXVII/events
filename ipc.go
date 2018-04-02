package main
/*
  Listens to a UNIX socket and forwards messages to the main function using a channel
*/

import (
  "net"
  "fmt"
  "bytes"
  
	"github.com/kjk/betterguid"
)

func Tasker(endpoint net.Listener, msgs chan Event) {
  for {
    conn, err := endpoint.Accept()
    if err != nil {
      panic(err)
    }
    go newTask(conn, betterguid.New(), msgs)
  }
}

func newTask(conn net.Conn, uid string, msgs chan Event){
  defer conn.Close()
  
  for {
    buf := make([]byte, 1500) // MTU size
  	nr, err := conn.Read(buf)
  
  	if err != nil {
  	  fmt.Println(err)
  		break
  	}
  
  	data := bytes.Split(bytes.Replace(buf[0:nr], []byte("\n"), []byte("\\n"), -1), []byte(": "))
    fmt.Println(data)
    
    msgs <- Event{string(data[0]),fmt.Sprintf("{\"%s\":\"%s\"}", uid, data[1])}
  }
}
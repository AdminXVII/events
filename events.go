package main

import (
  "log"
  "net/http"
  
  "github.com/AdminXVII/go-sse"
)

type Event struct {
  Name string
  Msg  string
}

func main() {
  // Create the server.
  s := sse.NewServer(&sse.Options{
    Headers: map[string]string {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Methods": "GET, OPTIONS",
      "Access-Control-Allow-Headers": "Keep-Alive,Cache-Control,Content-Type",
    },
  })

  defer s.Shutdown()
  
  unix, tcp := Listeners()
  
  // Create tasks channel
  msgs := make(chan Event, 100)
  
  go Tasker(unix, msgs)
  
  go func(){
    for msg := range msgs {
      s.SendMessage("", sse.NewMessage("", msg.Msg, msg.Name))
    }
  }()
  
  // Register with endpoint.
  http.Handle("/", s)
  
  // Launch server
  log.Fatal("Fatal error: ", http.Serve(tcp, nil))
}
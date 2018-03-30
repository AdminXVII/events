package sse
/*
  Server for SSE. Axed on channels
*/

import (
  "fmt"
  "net/http"
  "log"
)


type server struct {
  http.Handler
}

type client struct {
  events chan Event;
}

func NewServer(init func(*Client)) *Server {
  return &Server{}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
  // Set the headers related to event streaming.
  rw.Header().Set("Content-Type", "text/event-stream")
  rw.Header().Set("Cache-Control", "no-cache")
  rw.Header().Set("Connection", "keep-alive")
  rw.Header().Set("Access-Control-Allow-Methods", "GET")
  rw.Header().Set("Access-Control-Allow-Origin", "*")
  
  // Make sure that the writer supports flushing.
  flusher, ok := rw.(http.Flusher)
  
  if !ok {
      http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
      return
  }
  
  c := s.NewClient()
  s.Connect(c)
  s.DisconnectOn(c, rw.(http.CloseNotifier).CloseNotify())
	
	log.Println("Finished init phase")
}

func (s *Server) Connect(c *Client) {
  s.clients = append(s.clients, c)
}

func (s *Server) Broadcast(event Event) {
	log.Println("Client")
	for c := range s.clients {
    c <- event
	}
}

func (s *Server) NewClient() *Client{
  c := { make(chan Event) }
}
  
func (s *Server) DisconnectOn(c *Client, disconnect chan bool) {
  go func(){
  	<-disconnect // On client close
  	c.Die()
  }()
}

func (c *Client) Connect(events chan Event) {
  for event := range events {
    fmt.Fprintf(rw, "event: %s\ndata: %s\n\n", event.Name, event.Msg)
    flusher.Flush()
    log.Println("Client sent info")
	}
}
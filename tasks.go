package main

import (
  "log"
  "net/http"
  "time"
  
  "github.com/coreos/go-systemd/activation"
  "github.com/alexandrevicenzi/go-sse"
  
  "github.com/AdminXVII/tasks-go/ipc"
)

//type Event struct {
//  Name string
//  Msg  string
//}

func debugLog() {
  f, err := os.OpenFile("testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()
  
  log.SetOutput(f)
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
  
  listeners := Listeners()
  
  // Create tasks channel
  msgs := make(chan ipc.Event, 100)
  
  go ipc.Tasker(msgs)
  
  go func(){
    for msg := range msgs {
      s.SendMessage("", sse.NewMessage("", msg.Msg, msg.Name))
    }
  }()
  
  go func(){
    for true {
      msgs <- ipc.Event{"new","{\"asdasds\":\"sadasdjasvhdjashdjashvdjhsvdjhasvd\"}"}
      time.Sleep(5)
    }
  }()
  
  // Register with endpoint.
  http.Handle("/", s)
  
  // Launch server
  log.Fatal("Fatal error: ", http.ListenAndServe(listeners[0], nil))
}

func Listeners() [2]string{
  // Defaults
	out := [2]string{":9000","/tmp/tasks.sock"}
	
  listeners, err := activation.Listeners(true)
	if err != nil {
		panic(err)
	}

	if len(listeners) > 2 {
		panic("Unexpected number of socket activation fds")
	}
	
	log.Println(listeners)
	
	return out
}
package ipc
/*
  Listens to a UNIX socket and forwards messages to the main function using a channel
*/

type MsgType int

const (
  Add MsgType = iota
  Msg
  End
)

type task struct {
  name string
  uid  int
}

type Message struct {
  Uid int
  Msg string
  Type MsgType
}

type Event struct {
  Name string
  Msg  string
}

func Tasker(msgs chan Event) {
  
}
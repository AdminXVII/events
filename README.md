# Events

Multiplex events as server-sent event

## Installation

    go get github.com/AdminXVII/events
    cd $GOPATH/github.com/AdminXVII/events
    go install

## Systemd integration

    sudo cp events /usr/local/bin
    sudo cp systemd/* /etc/systemd/user
    systemd --user daemon-reload
    systemd --user enable events

## Format

On the local side, communication is done using sequential packets through a Unix socket. Message are composed of an event name, followed by a colon and a space, than the message. i.e. `myevent: my fantastic message`

On the remote side, SSE protocol is used. The data consist of a dictionary mapping connections UIDs to the message. JSON data is not decoded and simply transferred as a string.

    event: myevent
    data: my fantastic message

## Example usage

Open the *Events* server, go to chezxavier.ga/tasks on the same computer, enter `:9000` in the field.
Open a terminal and type `socat - /tmp/tasks.sock,TYPE=5` (or `socat - $XDG_RUNTIME_DIR/tasks.sock,TYPE=5` if using systemd)
Type `new: A sample task <return>` and watch the event being sent to the browser's window

## TODO

 - Add support for SSL
 - Command-line and config file
 - Packages for various distribution
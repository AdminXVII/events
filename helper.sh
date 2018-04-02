#! /bin/sh

reload() {
    systemctl --user stop events
    go build
    sudo cp events /usr/bin/events
    systemctl --user start events
}

status() {
    systemctl --user status tasks
}

connect() {
    socat - /run/user/1000/tasks.sock,type=5
}
package main

import (
  "testing"
  "log"
  "net/http"
)

func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    log.Println(Listeners())
}

func Benchmarkseparate
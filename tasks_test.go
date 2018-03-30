package main

import (
  "testing"
  "log"
)

func TestTimeConsuming(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping test in short mode.")
    }
    log.Println(Listeners())
}
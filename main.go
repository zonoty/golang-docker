package main

import (
  "fmt"
  "net/http"

  "github.com/zonoty/golang-docker/auth"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")
}

func main() {
  auth.ExecLog()
  http.HandleFunc("/", handler)
  http.ListenAndServe(":80", nil)
}

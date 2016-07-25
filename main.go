package main

import (
  "fmt"
  "os"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hello world")
}

func main()  {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}


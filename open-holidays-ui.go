package main

import (
  "html/template"
  "net/http"
  "fmt"
)


func main() {
  const port = "3001"
  const portAddr = ":" + port
  http.HandleFunc("/request-leave", requestLeaveHandler)
  fmt.Println("Listening on port", port)
  http.ListenAndServe(portAddr, nil)
}

func requestLeaveHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("GOT HERE %v", r.Method)
  switch r.Method {
  case "GET":
    fmt.Println("GOT TO GET")
    t, err := template.ParseFiles("views/request-leave.html")
    if err != nil {
      fmt.Fprintf(w, "Sad times, we messed up and this page won't load. Please try later")
    }
    t.Execute(w, nil)
  case "POST":
    fmt.Println("GOT TO POST")
    t, err := template.ParseFiles("views/request-submitted.html")
    if err != nil {
      fmt.Fprintf(w, "Sad times, we messed up and this page won't load. Please try later")
    }
    t.Execute(w, nil)
  }
}

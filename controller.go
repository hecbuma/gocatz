package main

import (
  "fmt"
  "net/http"

  // "github.com/gorilla/mux"
  "code.google.com/p/google-api-go-client/customsearch/v1"
)

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Catz Catz Catz!")
}

func Show(w http.ResponseWriter, r *http.Request) {
  // vars := mux.Vars(r)
  // height := vars["height"]
  // width := vars["width"]

  cat := Cat{url: "http://ecx.images-amazon.com/images/I/51sxD3in7mL._SY300_.jpg", width: 200, height: 300}
  fmt.Fprintf(w, "%q", cat.url)
}

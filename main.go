package main

import (
    "encoding/json"
    "log"
    "net/http"
    "math/rand"
    "strconv"
    "github.com/gorilla/mux"
)

// Book Struct (Model)

type Book struct {
    ID string `json:"id"`
    Isbn string `json:"isbn"`
    Title string `json:"title"`
    Author *Author `json:"author"`
}

// Author Struct
type Author struct {
    Firstname string `json:"firstname"`
    Lastname string `json:"lastname"`
}

func main() {
    r :=
}

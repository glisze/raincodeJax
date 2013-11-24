package hello

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/get", get)
    http.HandleFunc("/post", post)
    http.HandleFunc("/put", put)
    http.HandleFunc("/delete", delete)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi, Lo, this is the plat formed app-engine!\n")
}

func get(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi! What do you want?\n")
}

func post(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi! Thank You, very, very much! You will hear from us very very; soon. Promise!\n")
}

func put(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi! Thanks! But what am I going to do with it?\n")
}

func delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hi! That was restful. Again!\n")
}

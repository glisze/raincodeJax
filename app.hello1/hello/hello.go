package hello

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/post", post)
	http.HandleFunc("/put", put)
	http.HandleFunc("/delete", delete)
}

func get(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
    fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Hello</title></head><body><h1>Hello,</h1><br/>%v!<br/>What's up?</body></html>", u)
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


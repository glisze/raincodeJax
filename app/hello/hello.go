package hello
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
 
    "appengine"
    "appengine/user"
    "github.com/rCooDr/onemath"
)

type Page struct {
    Title string
    Body  []byte
}

func (p *Page) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
 
func init() {
    http.HandleFunc("/get", get)
    http.HandleFunc("/put", put)
    http.HandleFunc("/post", post)
    http.HandleFunc("/delete", delete)
    http.HandleFunc("/Sqrt", handler)
}
 
func handler(w http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf(w, "<h1>Hello,</h1>\n<p>%v!</p><pre>Sqrt(%v) = %v</pre>", u, 4, onemath.Sqrt(4))
}
 
func get(w http.ResponseWriter, r *http.Request) {

    p2, _ := loadPage("TestPage")

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
    fmt.Fprintf(w, "<pre>Hello, %v!</pre><hr/>", u)
    fmt.Println(string(p2.Body))
}
 
func put(w http.ResponseWriter, r *http.Request) {

    p1 := &Page{Title: "TestPage", Body: []byte("<h1>This is a sample Page.</h1>\n<p>squeezed into the app</p>")}
    p1.save()

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
    fmt.Fprintf(w, "<h1>put</h1> <p>%v!</p>", u)
}
 
func post(w http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf(w, "<h1>post</h1><p>%v!</p>", u)
}
 
func delete(w http.ResponseWriter, r *http.Request) {
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
    fmt.Fprintf(w, "<h1>delete</h1> <p>%v!</p>", u)
~/qnjp-uQ-Vn6/app/hello$

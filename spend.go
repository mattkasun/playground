package main

import (
                "fmt"
                "log"
                "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
                if r.URL.Path != "/" {
                                http.Error(w, "404 Not Found\n" +  r.URL.Path, http.StatusNotFound)
                                return
                }
                switch r.Method {
                case "GET":
                                http.ServeFile(w, r, "form.html")
                case "POST":
                                if err := r.ParseForm(); err != nil {
                                                fmt.Fprintf(w, "ParseForm() err: %v", err)
                                                return
                                }
                                fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
                                name := r.FormValue("name")
                                address := r.FormValue("address")
                                fmt.Fprintf(w, "Name = %s\n", name)
                                fmt.Fprintf(w, "Address = %s\n", address)
    fmt.Fprintf(w, "Button = %s\n", r.FormValue("OK"))
    fmt.Fprintf(w, "Request " , r)
                default:
                                fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
                }
}

func main() {
  fs := http.FileServer(http.Dir("stylesheet"))
  http.Handle("/stylesheet/", http.StripPrefix("/stylesheet/", fs))
                http.HandleFunc("/", helloHandler)
                log.Fatal(http.ListenAndServe(":8080", nil))
}

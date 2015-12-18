package hello

import (
    "fmt"
    "log"
    "net/http"
)

func init() {
    http.HandleFunc("/hello", helloworld)
    http.HandleFunc("/join", join)
    http.HandleFunc("/start", start)
    http.HandleFunc("/stop", stop)
}

func helloworld(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "Hello, world!")

    for key, value := range req.Form {
        log.Println(key)
        log.Println(value)
        //LOG: {"test": "that"}
    }
}

func join(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "joining")
}

func start(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "starting")
}

func stop(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "stopping!")
}
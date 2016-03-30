package main

import (
	"net/http"
	"log"
	"os"
	"fmt"
)

var info = log.New(os.Stdout,
	"INFO: ",
	log.Ldate|log.Ltime|log.Lshortfile)



func handler(w http.ResponseWriter, r *http.Request) {
	info.Println(r.RemoteAddr)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("0.0.0.0:23333", nil)
}
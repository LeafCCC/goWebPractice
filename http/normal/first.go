package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello,"+name)
}

func shoutHelloServer(w http.ResponseWriter, req *http.Request) {
	name := req.URL.Path[len("/shoutHello/"):]
	fmt.Fprintf(w, "Hello,"+strings.ToUpper(name))
}

func main() {
	http.HandleFunc("/hello/", helloServer)
	http.HandleFunc("/shoutHello/", shoutHelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

package main

import (
	"fmt"
	"net/http"
)

func handleFoo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling a Request from ", r.Host)
	fmt.Fprintf(w, "foo requested")
}

func startBasicServer() {

	//set handlers
	http.HandleFunc("/foo", handleFoo)

	fmt.Println("Start http server on port :80")
	http.ListenAndServe(":80", nil)

}

func main() {
	startBasicServer()
}

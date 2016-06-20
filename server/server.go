package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func simpleHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}
func StartServer() {

	r := mux.NewRouter()
	//post request
	r.HandleFunc("/sayhi", simpleHandler)

	http.Handle("/", r)
	http.ListenAndServe(":9091", nil)
}
func main() {
	StartServer()
}

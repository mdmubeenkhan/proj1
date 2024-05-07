package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(response http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Println(response, "ParseForm() error = ", err)
		return
	}
	fmt.Fprintf(response, "Form post request is successful")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "name = %s\n", name)
	fmt.Fprintf(response, "address = %s\n", address)
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(response, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(response, "hello from go app!")
}

func main() {

	fmt.Println("server started")
	fileServer := http.FileServer((http.Dir("./static")))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

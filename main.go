package main

import (
	"fmt"
	"net/http"
)

func main1() {
	fmt.Print("Hello, World!")
}

func main() {
	//	 set a web server
	//	 listen to port 8080
	//	 handle request
	//	 return response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World this is web Service in GO!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})

	http.ListenAndServe(":3000", nil)
}

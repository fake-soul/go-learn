package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main1() {
	fmt.Print("Hello, World!")
}

func main2() {
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

func main() {
	level := flag.String("level", "INFO", "level of logging")
	flag.Parse()

	f, err := os.Open("logs.txt")

	if err != nil {
		fmt.Println("Error opening file")
		log.Fatal(err)
		os.Exit(1)
	}

	defer f.Close()

	bufReader := bufio.NewReader(f)

	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}

//go run . -level DEBUG

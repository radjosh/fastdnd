package main

import (
	"fmt"
	"errors"
	"io"
	"net/http"
	"os"
)

import "dataProcessors"

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	// fmt.Print("")
	// fmt.Println(dataProcessors.Monsters["Dao"])
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

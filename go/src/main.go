package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"errors"
	// "io"
	"net/http"
	"os"
)

import "dataProcessors"

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got / request\n")
// 	io.WriteString(w, "This is my website!\n")
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

func monsterPtrToJSON(m *dataProcessors.Monster) {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(m)
	fmt.Println("\nUsing Encoder:\n" + buffer.String())
	// return buffer.String()
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	monsterName := r.URL.Query().Get("name")
	monsterPtr := new(dataProcessors.Monster)
	*monsterPtr = dataProcessors.Monsters[monsterName]
	monsterPtrToJSON(monsterPtr)
	// var payload = monsterPtrToJSON(monsterPtr)
	// io.WriteString(w, payload)
	// io.WriteString(w, payload)
	// io.WriteString(w, monsterPtr.Alignment)
}

func main() {
	// fmt.Print("")
	// fmt.Println(dataProcessors.Monsters["Dao"])
	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)
	http.HandleFunc("/monster", getMonster)
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

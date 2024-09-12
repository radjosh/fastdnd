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

func monsterPtrToJSON(m *dataProcessors.Monster) []byte {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(m)
	return buffer.Bytes()
}

func getMonster(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	monsterName := r.URL.Query().Get("name")
	monsterPtr := new(dataProcessors.Monster)
	*monsterPtr = dataProcessors.Monsters[monsterName]
	payload := monsterPtrToJSON(monsterPtr)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func main() {
	http.HandleFunc("/monster", getMonster)
	err := http.ListenAndServe(":3000", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

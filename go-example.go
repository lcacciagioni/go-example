package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "9000"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello, World!\n")
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		log.Printf("Warning, PORT not set. Defaulting to %+v\n", DEFAULT_PORT)
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Printf("ListenAndServe: ", err)
	}
}

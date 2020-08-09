package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Dummy server, does nothing!")
	})
	log.Println("Server started listening on port 1234")
	log.Fatal(http.ListenAndServe(":1234", nil))
}

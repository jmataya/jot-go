package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Thanks for the %s!", r.Method)
	})

	var err = http.ListenAndServe(":3000", nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		fmt.Printf("Listening on 3000")
	}
}

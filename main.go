package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Codespaces!!!")
	})

	log.Fatalln(http.ListenAndServe(":3000", nil))
}

package main

import (
	"fmt"
	"math/rand/v2"
	"log"
	"net/http"
)

func randomHandler(w http.ResponseWriter, r *http.Request) {
	a := rand.IntN(6)
	fmt.Fprint(w, "Random number: ", a+1)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", randomHandler)

	if err := http.ListenAndServe(":8085", mux); err != nil{
		log.Fatal(err)
	}

}
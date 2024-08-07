package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		Calculate(w,r,"+")
	})
	router.HandleFunc("POST /sub", func(w http.ResponseWriter, r *http.Request) {
		Calculate(w,r,"-")
	})
	router.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		Calculate(w,r,"*")
	})
	router.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		Calculate(w,r,"/")
	})

	fmt.Println("Listing on 1920")
	if err := http.ListenAndServe("0.0.0.0:1290", router); err != nil {
		log.Fatalln("failed to start the server:", err)
	}
}

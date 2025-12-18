package main

import (
	"fmt"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	//atlas.GenerateData(n, radius, seed)
}

func main() {
	fmt.Println("Listening on 8080")
	http.HandleFunc("/points", request)
	http.ListenAndServe(":8080", nil)
}

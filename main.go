package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		log.Println("req", req)
		resp.WriteHeader(200)
		resp.Write([]byte(`OK`))
	})
	log.Println("starting...")
	http.ListenAndServe("0.0.0.0:8080", nil)
}

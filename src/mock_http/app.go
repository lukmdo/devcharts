package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.Handle("/", loggerMW(http.HandlerFunc(handler)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

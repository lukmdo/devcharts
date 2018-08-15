package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"os"
)

var (
	Version = "?"
	VersionUTC = "?"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok: %s", Version)
}

func init() {
	argHelp := flag.Bool("help", false, "Show this help")
	argVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *argHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *argVersion {
		fmt.Printf("Version: %s\n", Version)
		fmt.Printf("UTC Build Time: %s\n", VersionUTC)
		os.Exit(0)

	}
}


func main() {
	port := "8080"
	http.Handle("/", loggerMW(http.HandlerFunc(handler)))
	log.Printf("Starting on port:%s version:%s", port, Version)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

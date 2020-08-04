package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var mode string

var staticAssetsDir = os.Getenv("STATIC_ASSETS_DIR")
var templatesDir = os.Getenv("TEMPLATES_DIR")

func main() {

	modePtr := flag.String("mode", "", "which mode to run")

	portPtr := flag.String("port", "8082", "Which port to run")
	flag.Parse()
	fmt.Println("word:", *modePtr)
	fmt.Println("port:", *portPtr)

	if *modePtr == "dev" {
		mode = "dev"
	}
	staticHandler := http.FileServer(http.Dir(staticAssetsDir))

	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(":"+*portPtr, nil))

}

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

var mode string

func main() {

	modePtr := flag.String("mode", "", "which mode to run")

	portPtr := flag.String("port", "8082", "Which port to run")
	flag.Parse()
	fmt.Println("word:", *modePtr)
	fmt.Println("port:", *portPtr)

	if *modePtr == "dev" {
		mode = "dev"
	}
	staticHandler := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))
	http.HandleFunc("/", indexHandler)

	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		// systemd run
		f := os.NewFile(3, "from systemd")
		l, err := net.FileListener(f)
		if err != nil {
			log.Fatal(err)
		}
		http.Serve(l, nil)
	} else {
		// manual run
		//		log.Fatal(http.ListenAndServe(":8080", nil))
		log.Fatal(http.ListenAndServe(":"+*portPtr, nil))
	}

}

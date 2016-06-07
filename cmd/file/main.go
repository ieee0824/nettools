package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port   string
	docDir string
)

func main() {
	flag.StringVar(&port, "p", "8080", "http server port")
	flag.StringVar(&port, "port", "8080", "http server port")
	flag.StringVar(&docDir, "d", "./", "document directory")
	flag.StringVar(&docDir, "documentDirectory", "./", "document directory")
	flag.Parse()

	log.Fatalln(http.ListenAndServe(":"+port, http.FileServer(http.Dir(docDir))))
}

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	body = flag.String("body", "", "HTTP response body.")
	port = flag.String("port", ":8080", "HTTP service port (e.g., ':8080')")
	code = flag.Int("code", 200, "HTTP response status code.")
)

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(*code)
		log.Print(req.Method, " ", req.RequestURI)
		fmt.Fprint(w, *body)
	})

	http.Handle("/", mux)

	log.Fatal(http.ListenAndServe(*port, mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"willnorris.com/go/imageproxy"
)

func main() {

	// Set server address
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("No port provided for the imageproxy")
	}

	// Create proxy
	p := imageproxy.NewProxy(nil, nil)

	// Create whitelist
	if os.Getenv("WHITELIST") != "" {
		p.Whitelist = strings.Split(os.Getenv("WHITELIST"), ",")
	}

	// Create baseurl
	if os.Getenv("BASEURL") != "" {
		var err error
		p.DefaultBaseURL, err = url.Parse(os.Getenv("BASEURL"))
		if err != nil {
			log.Fatalf("error parsing baseURL: %v", err)
		}
	}

	p.ScaleUp = true

	server := &http.Server{
		Addr:    port,
		Handler: p,
	}

	fmt.Printf("imageproxy listening on " + port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", port)
	}
}

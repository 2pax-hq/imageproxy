package main

import (
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/codegangsta/negroni"

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

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.Wrap(p),
	)

	n.Run(":" + port)
}

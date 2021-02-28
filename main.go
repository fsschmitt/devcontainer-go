package main

import (
	"log"

	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func hello() string {
	log.Printf("[INFO] %v", quoteV3.GlassV3())
	log.Printf("[INFO] %v", quoteV3.Concurrency())
	log.Printf("[INFO] %v", quoteV3.GoV3())
	return quote.Hello()
}

func main() {
	hello()
}

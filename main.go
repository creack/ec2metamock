package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var (
		flagSet = flag.NewFlagSet("sink", flag.ExitOnError)
		ip      = flagSet.String("ip", os.Getenv("DOCKER_IP"), "local ipv4 to report for `/latest/meta-data/local-ipv4`")
		port    = flagSet.Int("port", 8080, "port to listen to")
	)
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Connection", "Close")
		fmt.Fprintf(w, "%s", *ip)
	})))
}

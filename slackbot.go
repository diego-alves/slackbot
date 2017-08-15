package main

import (
	"flag"
)

var host = flag.String("host", "localhost:8080", "host:port to the server")
var server = flag.Bool("s", false, "Enable's server mode")

func main() {
	flag.Parse()

	if *server {
		Serve()
	} else {
		Connect(*host)
	}
	
}
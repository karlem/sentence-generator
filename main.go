package main

import (
	"flag"

	"github.com/karelm/sentence-generator/api"
)

var mode = flag.String("mode", "server", "Run either as http server 'server' or stress test with 'stress'.")

func main() {
	flag.Parse()

	if *mode == "server" {
		s := api.NewAPI()
		s.Run()
	} else {
		stressFile()
	}
}

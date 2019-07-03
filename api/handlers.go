package api

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (a API) learnHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()

	// Processing by line will allow us to process massive files.
	// We don't have to wait for the whole file to be processed so generation
	// will work almost instantly concurrently.
	scanner := bufio.NewScanner(r.Body)
	for scanner.Scan() {
		t := scanner.Text()

		// No need to try process
		if len(t) == 0 {
			continue
		}

		a.gen.Learn(t)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("/learn - Error occured during learning: %s", err)
		w.WriteHeader(500)
		fmt.Fprintln(w, "Error occured during learning:", err)
		return
	}

	w.WriteHeader(202)
	fmt.Fprintln(w, "Done")
	log.Printf("/learn - finished after: %s", time.Now().Sub(t))
}

func (a API) generateHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintln(w, a.gen.Generate())
	log.Printf("/generate - finished after: %s", time.Now().Sub(t))
}

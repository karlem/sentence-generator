package api

import (
	"log"
	"net/http"

	"github.com/karelm/sentence-generator/generator"
)

type API struct {
	gen generator.Generator
}

// NewAPI - return new instace of API
func NewAPI() API {
	return API{generator.NewGenerator()}
}

// Run - starts api server. Blocking.
func (a API) Run() {
	http.HandleFunc("/generate", a.generateHandler)
	http.HandleFunc("/learn", a.learnHandler)

	log.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

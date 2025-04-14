package main

import (
	"log"
	"net/http"

	"github.com/cmgchess/gotodo/router"
)

func main() {
	r := router.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}

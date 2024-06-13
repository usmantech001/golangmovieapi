package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/usmantech001/mongoapi/router"
)

func main() {
	fmt.Println("Mongo Api")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4400", r))
	fmt.Println("listenin to port 8080")
}

package main

import (
	"log"
	"net/http"
	routers "user-crud/router"
)

func main() {
	log.Println("main calling..")
	r := routers.Router()

	log.Println("server Started..")
	// starting our server on localhost:8080 port.
	log.Fatal(http.ListenAndServe(":8080", r))

}

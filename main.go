package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/someshkurrey/go-movie-demo/router"
)

func main() {

	fmt.Println("Nice Yar mongo db")
	r := router.Router()
	fmt.Println("server is staring...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server is listing on 4000 ...")

}

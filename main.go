package main

import (
	"apiwithdb/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Nice Yar mongo db")
	r := router.Router()
	fmt.Println("server is staring...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("server is listing on 4000 ...")

}

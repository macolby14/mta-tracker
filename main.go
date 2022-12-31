package main

import (
	"io"
	"log"
	"net/http"
)

func main(){
	log.Println("Starting main")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatalln(err)
	}
	

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Print(sb)
}
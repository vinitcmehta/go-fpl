package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")

	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s\n", data)
}

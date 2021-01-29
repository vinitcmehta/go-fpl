package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world!")

	//TODO convert json response of api into golang object
	res, err := http.Get("https://draft.premierleague.com/api/bootstrap-static")

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	//fmt.Printf("%s\n", data)
}

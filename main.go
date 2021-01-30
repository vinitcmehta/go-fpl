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
	//TODO inspect network tab to find APIs
	//TODO find top in form players based off last 3 weeks who aren't picked
	//TODO authenticate to my team
	//TODO get list of draft apis
	//https://draft.premierleague.com/api/bootstrap-static     players info
	//https://draft.premierleague.com/api/bootstrap-dynamic    user specific details of draft league
	//https://draft.premierleague.com/api/draft/44783/choices    drafted players in a league with given ID

	res, err := http.Get("https://draft.premierleague.com/api/bootstrap-static")

	if err != nil {
		log.Fatal(err)
	}

	data, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()

	fmt.Printf("%s\n", data)
}

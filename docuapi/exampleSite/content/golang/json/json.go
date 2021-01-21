package main

import (
	"log"

	jsoniter "github.com/json-iterator/go"
)

type AppResponse struct {
	Data        interface{} "json:\"data\""
	To          []string    "json:\"to\""
	StoreName   string      "json:\"storeName\""
	Concurrency string      "json:\"concurrency\""
}

//structAppResponse is the object describing the response from user code after a bindings event

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var resp AppResponse
	err := json.Unmarshal([]byte(`{"Data": "OK" }`), &resp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

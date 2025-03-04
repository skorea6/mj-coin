package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	port string = ":4000"
)

type URLDescriptions struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // 데이터가 없으면 없애줌
}

func documentation(rw http.ResponseWriter, req *http.Request) {
	data := []URLDescriptions{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	//b, err := json.Marshal(data)
	//utils.HandleErr(err)
	//fmt.Fprintf(rw, "%s", b)

	json.NewEncoder(rw).Encode(data)
}

func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))

	//explorer.Start()
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func genTimeJSON() string {
	type auxilaryStruct struct {
		Time string `json:"time"`
	}

	localTime := auxilaryStruct{string([]byte((time.Now()).Format(time.RFC3339)))}
	res, err := json.Marshal(localTime)

	if err != nil {
		fmt.Println(err)
	}

	return string(res)
}

func printTime(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, genTimeJSON())
}

func main() {
	http.HandleFunc("/time", printTime)
	log.Fatal(http.ListenAndServe(":8795", nil))
}

// simple server

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

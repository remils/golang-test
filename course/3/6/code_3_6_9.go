package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	GlobalId       int    `json:"global_id"`
	SystemObjectId string `json:"system_object_id"`
	SignatureDate  string `json:"signature_date"`
	Razdel         string
	Kod            string
	Name           string
	Idx            string
}

func main() {
	fo, err := os.Open("data-20190514T0100.json")
	if err != nil {
		panic(err)
	}

	defer fo.Close()

	var data []Data

	json.NewDecoder(fo).Decode(&data)

	var sum int = 0

	for _, item := range data {
		sum = sum + item.GlobalId
	}

	fmt.Println(sum)
}

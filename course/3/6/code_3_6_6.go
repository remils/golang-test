package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

type Student struct {
	Rating []int
}

type Data struct {
	Students []Student
}

type Average struct {
	Average float32
}

func main() {
	var count int = 0
	var data Data

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(stdin, &data); err != nil {
		panic(err)
	}

	for _, student := range data.Students {
		count = count + len(student.Rating)
	}

	var average Average = Average{float32(count) / float32(len(data.Students))}

	stdout, err := json.MarshalIndent(average, "", "    ")
	if err != nil {
		panic(err)
	}

	io.WriteString(os.Stdout, string(stdout))
}

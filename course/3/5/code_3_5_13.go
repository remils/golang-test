package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	arh, err := zip.OpenReader("task.zip")
	if err != nil {
		panic(err)
	}

	defer arh.Close()

	for _, f := range arh.File {
		rc, err := f.Open()
		if err != nil {
			panic(err)
		}

		r := csv.NewReader(rc)

		m, err := r.ReadAll()
		if err != nil {
			panic(err)
		}

		if len(m) == 10 && len(m[0]) == 10 {
			fmt.Println(m[4][2])
		}

		rc.Close()
	}
}

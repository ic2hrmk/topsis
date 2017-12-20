package main

import (
	"topsis/topsis"

	"io/ioutil"
	"encoding/json"
)

func readFile(fileName string) (report *topsis.Report, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &report)
	return
}
package main

import (
	"flag"
	"fmt"
	"os"
	"topsis/topsis"
)

var fileName = flag.String("file", "data.json", "File with formatted data (json format)")

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("FILE:    ", *fileName)

	fmt.Println("1. File read")
	report, err := readFile(*fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("2. Compute average expert's opinions")
	Y := topsis.GetAverageMarks(report)

	fmt.Println(Y)
}

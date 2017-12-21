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
	PrintFloatMatrix("Y", Y)

	fmt.Println("3. Calculate normalized weights and apply weights to averaged marks")
	W_n := topsis.GetNormalizedWeights(report.Weights)

	PrintFloatVector("Normalized weights Wn", W_n)

	Y_s := topsis.ApplyWeightedAverageMarks(Y, W_n)

	PrintFloatMatrix("Y'", Y_s)

	fmt.Println("4. Calculate best and worst Y points (Y+ & Y-)")
	Yplus, Yminus := topsis.GetReferencePoints(Y_s)

	PrintFloatVector("Y+", Yplus)
	PrintFloatVector("Y-", Yminus)

	fmt.Println("5. Calculate referenced distances")
	distances := topsis.GetDistancesToReferencePoints(Y_s, Yplus, Yminus)
	for _, d := range distances {
		fmt.Println(d)
	}
}

package topsis

import (
	"math"
	"strconv"
)

func GetAverageMarks(report *Report) (Y [][]float32) {
	expertNumber := len(report.Experts)

	X := report.Experts
	X_s := NewFloatCube(expertNumber, report.AlternativeNumber, report.CoefficientNumber)


	// Calculate X'iq(s) from Xiq(s)
	//	For each alternative
	for a := 0; a < report.AlternativeNumber; a++{
		//	Go trough all coefficients
		for k := 0; k < report.CoefficientNumber; k++ {
			//	Compute average expert mark

			expertsOpinions := NewFloatVector(expertNumber)

			for e := 0; e < expertNumber; e++ {
				expertsOpinions[e] = X[e][a][k]
			}

			summaryMark := Sum(expertsOpinions)

			for e := 0; e < expertNumber; e++ {
				X_s[e][a][k] = expertsOpinions[e] / float32(math.Sqrt(float64(summaryMark)))
			}
		}
	}

	// Average through all experts
	Y = NewFloatMatrix(report.AlternativeNumber, report.CoefficientNumber)

	//	All experts are equal in our case
	expertWeight := 1.0 / float32(expertNumber)
	for a := 0; a < report.AlternativeNumber; a++{
		//	Go trough all coefficients
		for k := 0; k < report.CoefficientNumber; k++ {
			shiftedExpertsOpinions := NewFloatVector(expertNumber)

			for e := 0; e < expertNumber; e++ {
				shiftedExpertsOpinions[e] = X_s[e][a][k]
			}

			shiftedSummaryMark := Sum(shiftedExpertsOpinions)

			Y[a][k] = expertWeight * shiftedSummaryMark
		}
	}

	return Y
}

func ApplyWeightedAverageMarks(Y [][]float32, W_n []float32) (Y_s [][]float32) {
	Y_s = NewFloatMatrix(len(Y), len(Y[0]))

	for i, w := range W_n {
		for a := range Y {
			Y_s[a][i] = Y[a][i] * w
		}
	}

	return
}

func GetNormalizedWeights(W []float32) (W_n []float32) {
	W_n = NewFloatVector(len(W))
	totalSum := Sum(W)

	for i := range W {
		W_n[i] = W[i] / totalSum
	}

	return
}

func GetReferencePoints(Y_s [][]float32) (Yplus, Yminus []float32) {
	Yplus = NewFloatVector(len(Y_s[0]))
	Yminus = NewFloatVector(len(Y_s[0]))

	for i := 0; i < len(Y_s[0]); i++ {
		column := GetColumn(Y_s, i)

		Yplus[i] = GetMax(column)
		Yminus[i] = GetMin(column)
	}

	return
}

func GetDistancesToReferencePoints(Y_s [][]float32, Yplus []float32, Yminus []float32) (distances []*Distance) {
	distances = make([]*Distance, len(Y_s))

	for i, y := range Y_s {
		dPlus := getEuclideanDistance(y, Yplus)
		dMinus := getEuclideanDistance(y, Yminus)

		alternativeName := "A" + strconv.Itoa(i+1)
		distances[i] = NewDistance(alternativeName , dPlus, dMinus)
	}

	return distances
}

func getEuclideanDistance(X []float32, Y[]float32) (distance float32) {
	for i := range X {
		distance += float32(math.Pow(float64(X[i] - Y[i]), float64(2)))
	}

	distance = float32(math.Sqrt(float64(distance)))

	return
}
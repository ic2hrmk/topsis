package topsis

import (
	"math"
)

func GetAverageMarks(report *Report) [][]float32 {
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
	Y := NewFloatMatrix(report.AlternativeNumber, report.CoefficientNumber)

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

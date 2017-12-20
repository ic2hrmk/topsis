package topsis

type Expert [][]float32

type Report struct {
	AlternativeNumber int        `json:"alternative_number"`
	CoefficientNumber int        `json:"coefficient_number"`

	Experts           []Expert	`json:"experts"`
	Weights           []float32	`json:"weights"`
}
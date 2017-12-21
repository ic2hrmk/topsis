package topsis

import "fmt"

type Expert [][]float32

type Report struct {
	AlternativeNumber int        `json:"alternative_number"`
	CoefficientNumber int        `json:"coefficient_number"`

	Experts           []Expert        `json:"experts"`
	Weights           []float32       `json:"weights"`
}

type Distance struct {
	AName  string
	DPlus  float32
	DMinus float32
	H      float32
}

func NewDistance(name string, dPlus float32, dMinus float32) *Distance {
	h := dMinus / (dPlus + dMinus)

	return &Distance{
		AName:name,
		DPlus: dPlus,
		DMinus: dMinus,
		H: h,
	}
}

func (d *Distance) String() string {
	return fmt.Sprintf("H(%s) = %3.3f",d.AName, d.H)
}
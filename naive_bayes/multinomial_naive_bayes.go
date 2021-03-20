package naive_bayes

import (
	"math"
)

const (
	CONSTANT = 1
)

type EvaluatorInterface interface {
	EvaluateInput(input interface{}) ([][]float64, error)
	GetTrainedData() map[string][][]float64
	GetDictionary() map[string]uint64
	GetSumVectorDataOfClass(class string) []float64
	GetSumDataOfClass(class string) float64
}

type MultinomialNaiveBa
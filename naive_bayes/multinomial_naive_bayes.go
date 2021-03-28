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

type MultinomialNaiveBayesConfig struct {
	Evaluator EvaluatorInterface
}

type MultinomialNaiveBayes struct {
	evaluator EvaluatorInterface
}

func NewMultinomialNaiveBayes(cfg MultinomialNaiveBayesConfig) MultinomialNaiveBayes {
	multinomialNaiveBayes := MultinomialNaiveBayes{
		evaluator: cfg.Evaluator,
	}

	return multinomialNaiveBayes
}

func (nb MultinomialNaiveBayes) Predict(inputs interface{}) ([]string, error) {
	var predicted []string
	probabilities, err := nb.PredictProbab
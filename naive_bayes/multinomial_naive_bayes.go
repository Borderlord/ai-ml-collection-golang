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
	probabilities, err := nb.PredictProbability(inputs)

	if err != nil {
		return nil, err
	}

	for _, prob := range probabilities {
		highestProb := float64(0)
		var selectedClass string
		for key, value := range prob {
			if highestProb < value {
				selectedClass = key
				highestProb = value
			}
		}
		predicted = append(predicted, selectedClass)
	}

	return predicted, nil
}

func (nb MultinomialNaiveBayes) PredictProb
package naive_bayes

import (
	"math"
)

const (
	CONSTANT = 1
)

type EvaluatorInterface interface {
	EvaluateInput(input interface{}) ([
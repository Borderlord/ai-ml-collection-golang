
package word_embedding

import (
	"errors"
	"fmt"
	"github.com/adrian3ka/go-learn-ai/helper"
	"math/rand"
	"time"
)

type LossFunctionType string
type ActivationFunctionType string
type OptimizerType string

const (
	OneHotEncodedNil = "One Hot Encoded Data is Nil"
	InvalidDataLearn = "Invalid Data Learn"

	CrossEntropy             LossFunctionType       = "cross_entropy"
	Softmax                  ActivationFunctionType = "softmax"
	GradientDescentOptimizer OptimizerType          = "gradient_descent_optimizer"
)

type Layer struct {
	Value  [][]float64
	Weight [][]float64
	Bias   float64
}

type WordEmbedding struct {
	dimension          uint64
	hiddenLayer        Layer
	outputLayer        Layer
	lossFunction       LossFunctionType
	activationFunction ActivationFunctionType
	optimizer          OptimizerType
	oneHotEncodedData  OneHotEncoderInterface
	learningRate       float64
}

type OneHotEncoderInterface interface {
	GetEncodedData() [][2][]uint64
}

type WordEmbeddingConfig struct {
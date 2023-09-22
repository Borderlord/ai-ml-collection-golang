
package tf_idf

import (
	"errors"
	"math"
)

const (
	UnequalDocumentLength = "UnequalDocumentLength"

	EuclideanSumSquare = "EuclideanSumSquare"
	EuclideanSum       = "EuclideanSum"
)

type CountVectorizer interface {
	GetDictionary() map[string]uint64
	VectorizedCounter() map[string][][]uint64
	Vectorize([]string) ([][]uint64, error)
}

type TermFrequencyInverseDocumentFrequency struct {
	smooth                   bool
	inverseDocumentFrequency []float64
	documentFrequency        []uint64
	data                     map[string][][]float64
	sumVectorDataPerClass    map[string][]float64
	sumDataPerClass          map[string]float64
	totalDocument            uint64
	normalizerType           string
	normalizer               []float64
	countVectorizer          CountVectorizer
}

type TermFrequencyInverseDocumentFrequencyConfig struct {
	Smooth          bool
	NormalizerType  string
	CountVectorizer CountVectorizer
}

func New(config TermFrequencyInverseDocumentFrequencyConfig) (TermFrequencyInverseDocumentFrequency, error) {
	tfidf := TermFrequencyInverseDocumentFrequency{
		smooth:          config.Smooth,
		normalizerType:  config.NormalizerType,
		countVectorizer: config.CountVectorizer,
	}

	tfidf.data = make(map[string][][]float64)
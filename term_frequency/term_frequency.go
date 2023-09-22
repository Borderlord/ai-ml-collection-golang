
package term_frequency

import (
	"strings"
)

type WordVectorizer interface {
	GetVectorizedWord() map[string]uint64
	Normalize(document string) (string, error)
}

type TermFrequency struct {
	data           map[string][][]uint64 //[corpus_name][][]word
	binary         bool
	wordVectorizer WordVectorizer
}

type TermFrequencyConfig struct {
	Binary         bool
	WordVectorizer WordVectorizer
}

package word_vectorizer

import (
	"regexp"
	"strings"
)

const (
	InvalidType = "Invalid Type"
)

type ArrayWordVectorizer struct {
	lower          bool
	data           map[string]uint64
	labelEncoded   [][2]uint64
	regexReplacers []RegexReplacer
}

type ArrayWordVectorizerConfig struct {
	Lower bool
}

func NewArrayWordVectorizer(vectorizer ArrayWordVectorizerConfig) *ArrayWordVectorizer {
	wv := 
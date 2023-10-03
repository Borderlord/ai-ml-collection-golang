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
	wv := ArrayWordVectorizer{
		lower: vectorizer.Lower,
		regexReplacers: []RegexReplacer{
			{Pattern: `[^a-zA-Z0-9 ]+`, Replacer: ``},
			{Pattern: `\s+`, Replacer: ` `},
		},
	}

	wv.data = make(map[string]uint64)
	return &wv
}

func (wv *ArrayWordVectorizer) Learn(arrayWord [][2]string) error {
	count := uint64(0
package word_vectorizer

import (
	"regexp"
	"strings"
)

type RegexReplacer struct {
	Pattern  string
	Replacer string
}

type WordVectorizer struct {
	lower           bool
	data            map[string]uint64 //[word]index
	cleanedCorpuses map[string][]string
	regexReplacers  []RegexReplacer
}

type WordVectorizerConfig struct {
	Lower bool
}

func New(vectorizer WordVectorizerConfig) WordVectorizer {
	wv := WordVectorizer{
		lower: vectorizer.Lower,
		regexReplacers: []RegexReplacer{
			{Pattern: `[^a-zA-Z0-9 ]+`, Replacer: ``},
			{Pattern: `\s+`, Replacer: ` `},
		},

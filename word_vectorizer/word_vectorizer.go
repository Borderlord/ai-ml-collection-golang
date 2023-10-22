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
	regexRepl
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
	lower           
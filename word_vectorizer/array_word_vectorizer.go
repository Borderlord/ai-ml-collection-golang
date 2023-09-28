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
	labelEncoded  
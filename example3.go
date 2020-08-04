package main

import (
	"fmt"
	"github.com/adrian3ka/go-learn-ai/one_hot_encoding"
	"github.com/adrian3ka/go-learn-ai/word_embedding"
	"github.com/adrian3ka/go-learn-ai/word_vectorizer"
)

func main() {
	vect := word_vectorizer.NewArrayWordVectorizer(word_vectorizer.ArrayWordVectorizerConfig{
		Lower: true,
	})

	err := vect.Learn([][2]string{
		{"king", "strong"},
		{"king", "man"},
		{"strong", "king"},
		{"strong", "man"},
		{"man", "king"},
		{"man", "strong"},
		{"queen", "wise"},
		{"queen", "woman"},
		{"wise", "queen"},
		{"wise", "woman"},
		{"woman", "queen"},
		{"woman", "wise"},
		{"boy", "young"},
		{"boy", "man"},
		{"young", "boy"},
		{"young", "man"},
		{"man", "boy"},
		{"man", "young"},
		{"girl", "young"},
		{"girl", "woman"},
		{"young", "girl"},
		{"young", "woman"},
		{"woman", "girl"},
		{"woman", "young"},
		{"prince", "young"},
		{"prince", "king"},
		{"young", "prince"},
		{"young", "king"},
		{"king", "prince"},
		{"king", "young"},
		{"princess", "young"},
		{"princess", "queen"},
		{"young", "princess"},
		{"young", "queen"},
		{"queen", "princess"},
		{"queen", "young"},
		{"man", "strong"},
		{"strong", "man"},
		{"woman", "pretty"},
		{"pretty", "woman"},
		{"prince", "boy"},
		{"prince", "king"},
		{"boy", "prince"},
		{"boy", "king"},
		{"king", "prince"},
		{"king", "boy"},
		{"princess", "girl"},
		{"princess", "queen"},
		{"girl", "princess"},
		{"girl", "queen"},
		{"queen", "princess"},
		{"queen", "girl"},
	})

	fmt.Println(err)

	fmt
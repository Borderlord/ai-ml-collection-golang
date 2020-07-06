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
		{"s
package main

import (
	"fmt"
	"github.com/adrian3ka/go-learn-ai/grammar_parser"
)

func main() {
	fmt.Println("==================================CHUNKING===========================================")
	taggedSentence := [][2]string{
		{"the", "DT"},
		{"little", "JJ"},
		{"yellow", "JJ"},
		
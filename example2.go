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
		{"dog", "NN"},
		{"barked", "VBD"},
		{"at", "IN"},
		{"the", "DT"},
		{"cat", "NN"},
	}

	gp, err := grammar_parser.NewRegexpParser(grammar_parser.RegexpParserConfig{
		Grammar: [][2]string{
			{"NP", 
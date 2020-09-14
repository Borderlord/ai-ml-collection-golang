
package grammar_parser

import (
	"errors"
	"github.com/adrian3ka/go-learn-ai/helper"
	"github.com/adrian3ka/go-learn-ai/nfa"
	"strings"
)

//AVAIABLE SYMBOL
//- CHINCKING : } {
//- CHUNGKING : { }

const (
	CannotCreateNFAFromGrammar    = "Cannot Create NFA From Grammar"
	InvalidGrammar                = "Invalid Grammar"
	NilState                      = "Nil State"
	ChinkingInitialStateMustBeNil = "Chinking Initial State Must Be Nil"
	OneOrMore                     = "+"
	NoneOrMore                    = "*"
	Optional                      = "?"
	OpeningTag                    = "<"
	ClosingTag                    = ">"

	Chunking = "Chunking"
	Chinking = "Chinking"
)

type BasicParser interface {
	Parse([][2]string) error
}

type NfaGrammar struct {
	Nfa            nfa.NFA
	Target         string
	AlreadyOnFinal bool
	Type           string
}

type RegexpParser struct {
	nfaGrammar []*NfaGrammar
}

type RegexpParserConfig struct {
	Grammar [][2]string
}

type HandleSymbolInput struct {
	PreviousState []*nfa.State
	NfaData       *nfa.NFA
	Tag           string
	IsFinal       bool
}

func handleNoneOrMore(input *HandleSymbolInput) (*nfa.NFA, []*nfa.State, error) {
	var newStates []*nfa.State
	var state1 *nfa.State
	var err error

	if input.NfaData == nil {
		input.NfaData, state1, err = nfa.NewNFA(input.Tag, input.IsFinal)

		if err != nil {
			return nil, nil, err
		}

	} else {

		state1, err = input.NfaData.AddState(&nfa.State{
			Name: input.Tag,
		}, input.IsFinal)

		if err != nil {
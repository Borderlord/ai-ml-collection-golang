
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
			return nil, nil, err
		}
	}

	newStates = append(newStates, state1)

	err = input.NfaData.AddTransition(state1.Index, input.Tag, *state1)

	if err != nil {
		return nil, nil, err
	}

	for idx, _ := range input.PreviousState {
		newStates = append(newStates, input.PreviousState[idx])
		err = input.NfaData.AddTransition(input.PreviousState[idx].Index, input.Tag, *state1)

		if err != nil {
			return nil, nil, err
		}

	}

	return input.NfaData, newStates, nil
}

func handleBasic(input *HandleSymbolInput) (*nfa.NFA, []*nfa.State, error) {
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
			return nil, nil, err
		}
	}

	newStates = append(newStates, state1)

	if err != nil {
		return nil, nil, err
	}

	for idx, _ := range input.PreviousState {
		err = input.NfaData.AddTransition(input.PreviousState[idx].Index, input.Tag, *state1)

		if err != nil {
			return nil, nil, err
		}
	}

	return input.NfaData, newStates, nil
}

func handleOptional(input *HandleSymbolInput) (*nfa.NFA, []*nfa.State, error) {
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
			return nil, nil, err
		}
	}

	newStates = append(newStates, state1)

	var tags []string
	tags = strings.Split(input.Tag, "|")

	for _, tag := range tags {
		err = input.NfaData.AddTransition(state1.Index, tag, *state1)
	}

	if err != nil {
		return nil, nil, err
	}

	for idx, _ := range input.PreviousState {
		newStates = append(newStates, input.PreviousState[idx])
		err = input.NfaData.AddTransition(input.PreviousState[idx].Index, input.Tag, *state1)

		if err != nil {
			return nil, nil, err
		}

	}

	return input.NfaData, newStates, nil
}

func handleOneOrMore(input *HandleSymbolInput) (*nfa.NFA, []*nfa.State, error) {
	var newStates []*nfa.State
	var state1 *nfa.State
	var err error

	if input.NfaData == nil {
		input.NfaData, state1, err = nfa.NewNFA(input.Tag, false)

		if err != nil {
			return nil, nil, err
		}

	} else {

		state1, err = input.NfaData.AddState(&nfa.State{
			Name: input.Tag,
		}, false)

		if err != nil {
			return nil, nil, err
		}
	}

	state2, err := input.NfaData.AddState(&nfa.State{
		Name: input.Tag,
	}, input.IsFinal)

	if err != nil {
		return nil, nil, err
	}

	if state1 == nil || state2 == nil {
		return nil, nil, errors.New(NilState)
	}

	for idx, _ := range input.PreviousState {
		err = input.NfaData.AddTransition(input.PreviousState[idx].Index, input.Tag, *state1)

		if err != nil {
			return nil, nil, err
		}
	}

	newStates = append(newStates, state2)

	err = input.NfaData.AddTransition(state1.Index, input.Tag, *state2)

	if err != nil {
		return nil, nil, err
	}

	err = input.NfaData.AddTransition(state2.Index, input.Tag, *state2)

	if err != nil {
		return nil, nil, err
	}

	return input.NfaData, newStates, nil
}

func handleInitialChunking(input *HandleSymbolInput) (*nfa.NFA, []*nfa.State, error) {
	var newStates []*nfa.State
	var state1 *nfa.State
	var err error

	if input.NfaData == nil {
		input.NfaData, state1, err = nfa.NewNFA(nfa.Negate+input.Tag, false)

		if err != nil {
			return nil, nil, err
		}

	} else {
		return nil, nil, errors.New(ChinkingInitialStateMustBeNil)
	}

	var tags []string
	tags = strings.Split(input.Tag, "|")

	for _, tag := range tags {
		err = input.NfaData.AddTransition(state1.Index, nfa.Negate+tag, *state1)
	}

	newStates = append(newStates, state1)

	return input.NfaData, newStates, nil
}
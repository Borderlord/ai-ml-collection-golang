package nfa

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	StateNotFound = "State Not Found"
	InvalidInput  = "Invalid Input"

	Negate = "!"
)

type transitionInput struct {
	srcStateIndex uint64
	input         string
}

type State struct {
	Name  string
	Index uint64
}

type destState map[State]bool

type NFA struct {
	initState    State
	currentState map[State]bool
	allStates    []State
	finalStates  []State
	transition   map[transitionInput]destState
	inputMap     map[string]bool
}

//New a new NFA
func NewNFA(initStateName string, isFinal bool) (*NFA, *State, error) {
	initState := State{
		Name:  initStateName,
		Index: 0,
	}
	retNFA := &NFA{
		transition: make(map[transitionInput]destState),
		inputMap:   make(map[string]bool),
		initState:  initState,
	}

	retNFA.currentState = make(map[State]bool)
	retNFA.currentState[initState] = true
	_, err := retNFA.AddState(&initState
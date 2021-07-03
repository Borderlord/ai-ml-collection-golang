package nfa

import (
	"testing"
)

func TestBasic(t *testing.T) {
	newNFA, state0, err := NewNFA("State 0", false)

	if err != nil {
		panic(err)
	}

	state1, err := newNFA.AddState(&State{
		Name: "State 1",
	}, false)

	if err != nil {
		panic(err)
	}

	state2, err := newNFA.AddState(&State{
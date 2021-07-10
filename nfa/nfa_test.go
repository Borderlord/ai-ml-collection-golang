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
		Name: "State 2",
	}, true)

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	err = newNFA.AddTransition(state0.Index, "a", *state1, *state2)

	if err != nil {
		panic(err)
	}

	e
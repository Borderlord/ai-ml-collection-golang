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
	_, err := retNFA.AddState(&initState, isFinal)

	if err != nil {
		return nil, nil, err
	}

	return retNFA, &initState, nil
}

func (d *NFA) GetCurrenteState() (map[State]bool, error) {
	return d.currentState, nil
}
func (d *NFA) GetAllState() ([]State, error) {
	return d.allStates, nil
}

//Add new state in this NFA
func (d *NFA) AddState(state *State, isFinal bool) (*State, error) {
	if state == nil || state.Name == "" {
		return nil, errors.New(InvalidInput)
	}

	currentIndex := uint64(len(d.allStates))

	state.Index = currentIndex

	d.allStates = append(d.allStates, *state)
	if isFinal {
		d.finalStates = append(d.finalStates, *state)
	}

	return state, nil
}

//Add new transition function into NFA
func (d *NFA) AddTransition(srcStateIndex uint64, input string, dstStateList ...State) error {
	find := false

	for _, v := range d.allStates {
		if v.Index == srcStateIndex {
			find = true
		}
	}

	if !find {
		return errors.New(StateNotFound)
	}

	if input == "" {
		return errors.New(InvalidInput)
	}

	//find input if exist in NFA input List
	if _, ok := d.inputMap[input]; !ok {
		//not exist, new input in this NFA
		d.inputMap[input] = true
	}

	dstMap := make(map[State]bool)
	for _, destState := range dstStateList {
		dstMap[destState] = true
	}

	targetTrans := transitionInput{srcStateIndex: srcStateIndex, input: input}
	d.transition[targetTrans] = dstMap

	return nil
}

func (d *NFA) Input(testInput string) ([]State, error) {
	updateCurrentState := make(map[State]bool)
	for current, _ := range d.currentState {
		intputTrans := transitionInput{srcStateIndex: current.Index, input: testInput}

		if valMap, ok := d.transition[intputTrans]; ok {
			for dst, _ := range valMap {
				updateCurrentState[dst] = true
			}
		} else {
			//dead state, remove in current state
			//do nothin
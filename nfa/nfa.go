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
	Index ui
package helper

import (
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func GetStringInBetween(str string, start string, end string) (result *string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, en
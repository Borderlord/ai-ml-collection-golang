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
	e := strings.Index(str, end)

	if s == -1 || e == -1 {
		return nil
	}

	ret := str[s:e]
	return &ret
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsAlphaNumeric(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return re.MatchString(s)
}

func IsAlphaUnderscore(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z_]*$")
	return re.MatchString(s)
}

func IsStringEqual(text string, characters []string) bool {
	for _, character := range characters {
		if text == character {
			return true
		}
	}
	return false
}

func LastSplit(text string, rune string) []string {
	result := strings.Split(text, rune)

	if len(result) < 2 {
		return []string{
			text,
		}
	}

	return []string{
		strings.Join(result[0:len(result)-1], rune),
		result[len(result)-1],
	}
}
func CalculateRecall(trueSlice interface{}, predictedSlice interface{}) float64 {
	trueReflection := reflect.ValueOf(trueSlice)
	if trueReflection.Kind() != reflect.Slice {
		panic("Interf
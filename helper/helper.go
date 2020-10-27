package helper

import (
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

func GetStringInBetween(str string, start string, end string) (re
package generator

import (
	"regexp"
	"strings"
)

var (
	padPunctuationReg = regexp.MustCompile("([^0-9])([.,!?])([^0-9])")
	punctuationReg    = regexp.MustCompile("[.,!?]")
)

func cleanInput(input string) []string {
	input = regexp.MustCompile("[()]").ReplaceAllString(input, "")
	input = regexp.MustCompile("([.-])+").ReplaceAllString(input, "$1")
	input = padPunctuationReg.ReplaceAllString(input, "$1 $2 $3")
	input = strings.NewReplacer(".", " . ", "\"", "").Replace(input)
	return strings.Fields(input)
}

func cleanOutput(output string) string {
	return regexp.MustCompile("\\s+([.,!?])\\s*").ReplaceAllString(output, "$1 ")
}

func sumFrequencies(m map[string]int) int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}

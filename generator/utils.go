package generator

import (
	"regexp"
	"strings"
)

var (
	padPunctuationReg                = regexp.MustCompile("([^0-9])([.,!?])([^0-9])")
	punctuationReg                   = regexp.MustCompile("[.,!?]")
	collapsePunctuationReg           = regexp.MustCompile("([,.-])+")
	collapseBracesReg                = regexp.MustCompile("[(){}\\[\\]]")
	collapseSpaceAfterPunctuationReg = regexp.MustCompile("\\s+([.,!?])\\s*")
)

func cleanInput(input string) []string {
	input = collapseBracesReg.ReplaceAllString(input, "")
	input = collapsePunctuationReg.ReplaceAllString(input, "$1")
	input = padPunctuationReg.ReplaceAllString(input, "$1 $2 $3")
	input = strings.NewReplacer(".", " . ", "\"", "").Replace(input)
	return strings.Fields(input)
}

func cleanOutput(output string) string {
	return collapseSpaceAfterPunctuationReg.ReplaceAllString(output, "$1 ")
}

func sumFrequencies(m map[string]int) int {
	var sum int
	for _, val := range m {
		sum += val
	}
	return sum
}

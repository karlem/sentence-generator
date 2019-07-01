package generator

import (
	"reflect"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestCreateTrigrams(t *testing.T) {
	corpus := "To be, or not to be, that is the question"
	expect := []trigram{
		{"To", "be", ","},
		{"be", ",", "or"},
		{",", "or", "not"},
		{"or", "not", "to"},
		{"not", "to", "be"},
		{"to", "be", ","},
		{"be", ",", "that"},
		{",", "that", "is"},
		{"that", "is", "the"},
		{"is", "the", "question"},
	}

	tokens := cleanInput(corpus)
	trigrams := createTrigrams(tokens)

	if reflect.DeepEqual(trigrams, expect) == false {
		t.Errorf("createTrigrams was incorrect, got: %v, want: %v.", trigrams, expect)
	}
}

// func TestGenerateText(t *testing.T) {
// 	input := "To be, or not to be, that is the question"
// 	// expect := [][3]string{{"To", "be", "or"}, {"be", "or", "not"}, {"or", "not", "to"}, {"not", "to", "be"}, {"to", "be", "that"}, {"be", "that", "is"}, {"that", "is", "the"}, {"is", "the", "question"}}

// 	g := NewGenerator()
// 	g.Learn(input)
// 	g.Generate("", 5)

// 	// reflect.DeepEqual(g.trigrams, expect)
// }

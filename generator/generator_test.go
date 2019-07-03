package generator

import (
	"reflect"
	"testing"
)

func TestCreateTrigrams(t *testing.T) {
	corpus := "To be, or not to be, that is the question."
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
		{"the", "question", "."},
	}

	tokens := cleanInput(corpus)
	trigrams := createTrigrams(tokens)

	if reflect.DeepEqual(trigrams, expect) == false {
		t.Errorf("createTrigrams was incorrect, got: %v, want: %v.", trigrams, expect)
	}
}

func TestLearnAddTrigramsFrequencies(t *testing.T) {
	input := "To be, or not to be, that is the question."
	expectPairs := []keyPair{
		keyPair{"To", "be"},
		keyPair{"be", ","},
		keyPair{",", "or"},
		keyPair{"or", "not"},
		keyPair{"not", "to"},
		keyPair{"to", "be"},
		keyPair{",", "that"},
		keyPair{"that", "is"},
		keyPair{"is", "the"},
		keyPair{"the", "question"},
	}
	expectFreq := map[keyPair]map[string]int{
		keyPair{",", "or"}:         {"not": 1},
		keyPair{",", "that"}:       {"is": 1},
		keyPair{"To", "be"}:        {",": 1},
		keyPair{"be", ","}:         {"or": 1, "that": 1},
		keyPair{"is", "the"}:       {"question": 1},
		keyPair{"not", "to"}:       {"be": 1},
		keyPair{"or", "not"}:       {"to": 1},
		keyPair{"that", "is"}:      {"the": 1},
		keyPair{"to", "be"}:        {",": 1},
		keyPair{"the", "question"}: {".": 1},
	}

	g := NewGenerator()
	g.Learn(input)

	// I don't like to use type assertion.
	// But this is just for unit test so should be fine.
	gt, ok := g.(*generator)
	if !ok {
		t.Error("Generator interface is expected to have generator type")
	}

	if !reflect.DeepEqual(gt.trigramsFreq, expectFreq) {
		t.Errorf("trigramsFreq was incorrect: \n got: %v \nwant: %v.", gt.trigramsFreq, expectFreq)
	}

	if !reflect.DeepEqual(gt.trigramsFreq, expectFreq) {
		t.Errorf("expectPairs was incorrect: \n got: %v \nwant: %v.", gt.pairs, expectPairs)
	}
}

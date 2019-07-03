package generator

import (
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"time"
)

const separator = " "

var (
	trigramsFreqMutex = sync.RWMutex{}
	endOfSentenceReg  = regexp.MustCompile("[.!?]")
)

// Generator - thread safe random text generator.
// Generate - generates randomly N number of sentences based on corpus. Return empty string if corpus has not been process.
// Learn - upload and process corpus
type Generator interface {
	Learn(string)
	Generate() string
}

type keyPair struct {
	f, s string
}

// Pairs are used for quick random lookup
type generator struct {
	pairs        []keyPair
	trigramsFreq map[keyPair]map[string]int
}

// NewGenerator - returns new instance of Generator
func NewGenerator() Generator {
	return &generator{
		pairs:        []keyPair{},
		trigramsFreq: make(map[keyPair]map[string]int),
	}
}

func (g *generator) getRandomKey() keyPair {
	trigramsFreqMutex.RLock()
	k := g.pairs[rand.Intn(len(g.pairs))]
	trigramsFreqMutex.RUnlock()

	// Find key which does not start with punctuation
	if punctuationReg.MatchString(k.f) == true {
		return g.getRandomKey()
	}

	return k
}

func (g *generator) nextWord(key keyPair) string {
	// Get a random key if previous can not be found
	// This can happen when we reach last trigram
	trigramsFreqMutex.RLock()
	choices, ok := g.trigramsFreq[key]
	trigramsFreqMutex.RUnlock()

	if !ok {
		key := g.getRandomKey()
		trigramsFreqMutex.RLock()
		choices = g.trigramsFreq[key]
		trigramsFreqMutex.RUnlock()
	}

	trigramsFreqMutex.RLock()
	totalWeight := sumFrequencies(choices)
	trigramsFreqMutex.RUnlock()

	r := rand.Intn(totalWeight)
	upto := 0

	var nextWord string

	// Weighted random selection
	// https://medium.com/@peterkellyonline/weighted-random-selection-3ff222917eb6
	trigramsFreqMutex.RLock()
	for choice, weight := range choices {
		upto += weight
		if upto > r {
			nextWord = choice
			break
		}
	}
	trigramsFreqMutex.RUnlock()

	return nextWord
}

// Generate map of frequinces from trigrams
// eg: from 'to be or' -> map[keyPair]map[string]int{keyPair{f:"to" s:"be"}: {"or": 1}}
// The reason for keyPair is because it allows quick lookups without a need of parsing (in case of text)
func (g *generator) addTrigramsFrequencies(trigrams []trigram) {
	for _, trigram := range trigrams {
		key := keyPair{trigram[0], trigram[1]}
		lastToken := trigram[2]

		trigramsFreqMutex.RLock()
		_, ok := g.trigramsFreq[key]
		trigramsFreqMutex.RUnlock()

		if !ok {
			trigramsFreqMutex.Lock()
			g.trigramsFreq[key] = map[string]int{lastToken: 1}
			g.pairs = append(g.pairs, key)
			trigramsFreqMutex.Unlock()
			continue
		}

		trigramsFreqMutex.Lock()
		g.trigramsFreq[key][lastToken]++
		trigramsFreqMutex.Unlock()
	}
}

func (g *generator) Generate() string {
	rand.Seed(time.Now().UnixNano())

	trigramsFreqMutex.RLock()
	len := len(g.trigramsFreq)
	trigramsFreqMutex.RUnlock()

	// No data to generate text from
	if len == 0 {
		return ""
	}

	currentKey := g.getRandomKey()
	var randText strings.Builder

	// Add first two words from key pair
	randText.WriteString(strings.Title(currentKey.f))
	randText.WriteString(separator)
	randText.WriteString(currentKey.s)

	// Generates at least 2 sentences
	numberOfSenteces := rand.Intn(50) + 2
	generatedSentences := 0
	for generatedSentences < numberOfSenteces {
		nextWord := g.nextWord(currentKey)

		randText.WriteString(separator)
		randText.WriteString(nextWord)

		currentKey = keyPair{currentKey.s, nextWord}

		// If text ends with [.!?] count as new sentence
		if endOfSentenceReg.MatchString(currentKey.s) == true {
			generatedSentences++
		}

		// In case the text does not contain any punctuation.
		// Maximum number of sentences is cca 5000 bytes
		if generatedSentences == 0 && randText.Len() > 5000 {
			break
		}
	}

	return cleanOutput(randText.String())
}

func (g *generator) Learn(corpus string) {
	trigrams := createTrigrams(cleanInput(corpus))
	g.addTrigramsFrequencies(trigrams)
}

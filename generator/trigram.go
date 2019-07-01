package generator

const numberOfTokens = 3

type trigram [3]string

func createTrigrams(tokens []string) []trigram {
	trigrams := []trigram{}

	for i := 0; i < len(tokens)-numberOfTokens+1; i++ {
		trigrams = append(trigrams, trigram{tokens[i], tokens[i+1], tokens[i+2]})
	}

	return trigrams
}

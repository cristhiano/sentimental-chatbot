package nlp

import (
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

func LexiconScore(text string) float64 {
	parsedtext := sentitext.Parse(text, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)

	return sentiment.Compound
}

package conversation

import (
	"github.com/cristhiano/sentimental-chatbot/nlp"
)

type NLPResolver string

var (
	NLPResolverLexicon NLPResolver = "lexicon"
	NLPResolverML      NLPResolver = "machine-learning"
)

type Interaction struct {
	Statement        string
	NLPResolver      NLPResolver
	PositiveAnswer   *Interaction
	NegativeAnswer   *Interaction
	AmbivalentAnswer *Interaction
	ExtractFunction  func(string)
}

func (i *Interaction) Resolve(input string) *Interaction {
	if i.ExtractFunction != nil {
		i.ExtractFunction(input)
	}

	switch i.NLPResolver {
	case NLPResolverLexicon:
		return i.ResolveWithLexicon(input)
	case NLPResolverML:
		return i.ResolveWithML(input)
	default:
		return i.ResolveWithML(input)
	}
}

func (i *Interaction) ResolveWithLexicon(input string) *Interaction {
	score := nlp.LexiconScore(input)

	if score < -0.1 {
		return i.NegativeAnswer
	}

	if score > 0.1 {
		return i.PositiveAnswer
	}

	return i.AmbivalentAnswer
}

func (i *Interaction) ResolveWithML(input string) *Interaction {
	score := nlp.MachineLearningScore(input)

	switch score {
	case 0:
		return i.NegativeAnswer
	case 1:
		return i.PositiveAnswer
	default:
		return i.AmbivalentAnswer
	}
}

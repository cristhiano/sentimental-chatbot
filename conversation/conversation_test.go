package conversation

import (
	"log"
	"testing"
)

func TestStart(t *testing.T) {
	petsInteraction := Interaction{
		Statement:   "Do you have any pets?",
		NLPResolver: NLPResolverLexicon,
		PositiveAnswer: &Interaction{
			Statement:   "What's their name(s)?",
			NLPResolver: NLPResolverLexicon,
			PositiveAnswer: &Interaction{
				Statement: "That's awesome!",
			},
			NegativeAnswer: &Interaction{
				Statement: "Chill, just asking",
			},
			AmbivalentAnswer: &Interaction{
				Statement: "Ok then",
			},
		},
		NegativeAnswer: &Interaction{
			Statement:   "Don't like animals?",
			NLPResolver: NLPResolverLexicon,
			PositiveAnswer: &Interaction{
				Statement: "You should get a pet then!",
			},
			NegativeAnswer: &Interaction{
				Statement: "I see",
			},
			AmbivalentAnswer: &Interaction{
				Statement: "I see",
			},
		},
	}

	tests := []struct {
		name        string
		interaction Interaction
		answers     []string
	}{
		{
			name: "Checking",
			interaction: Interaction{
				Statement:   "Hello, wanna chat?",
				NLPResolver: NLPResolverLexicon,
				PositiveAnswer: &Interaction{
					Statement: "How are you?",
					PositiveAnswer: &Interaction{
						Statement: "Great!",
					},
					NegativeAnswer: &Interaction{
						Statement: "I'm sorry to hear that",
					},
				},
				NegativeAnswer: &Interaction{
					Statement: "Okay, have a great day!",
				},
			},
			answers: []string{
				"yeah",
				"I'm fine",
			},
		},
		{
			name:        "Pets no",
			interaction: petsInteraction,
			answers: []string{
				"no",
				"not so much",
			},
		},
		{
			name:        "Pets yes",
			interaction: petsInteraction,
			answers: []string{
				"yes",
				"mind your own business",
			},
		},
		{
			name:        "Pets yes 2",
			interaction: petsInteraction,
			answers: []string{
				"yes",
				"why the hell do you want to know?",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conversation{
				Interaction: tt.interaction,
			}

			human := make(chan string)
			robot := c.Start(human)

			var i int
			for statement := range robot {
				log.Print(statement)

				if i == len(tt.answers) {
					break
				}

				log.Print(tt.answers[i])
				human <- tt.answers[i]

				i++
			}
		})
	}
}

package scripts

import (
	"fmt"
	"regexp"
	"strconv"

	"log"

	"github.com/cristhiano/sentimental-chatbot/conversation"
)

type Review struct {
	Product     string
	UserID      string
	Rating      int
	ValidRating bool
}

func (r Review) String() string {
	return fmt.Sprintf(
		"User id: %s Product: \"%s\" Rating: %d Valid: %v",
		r.UserID, r.Product, r.Rating, r.ValidRating)
}

type ReviewConversation struct {
	conversation.Conversation
	*Review
}

func NewReviewConversation(product, userID string) *ReviewConversation {
	review := &Review{
		Product: product,
		UserID:  userID,
	}

	rateInteraction := &conversation.Interaction{
		Statement: `Fantastic! On a scale of 1-5, how would you rate the iPhone 13?`,
		PositiveAnswer: &conversation.Interaction{
			Statement: `Thank you for sharing your feedback! If you have any more thoughts or need assistance with anything else, feel free to reach out!`,
		},
		NegativeAnswer: &conversation.Interaction{
			Statement: `Thank you for sharing your feedback! If you have any more thoughts or need assistance with anything else, feel free to reach out!`,
		},
		ExtractFunction: func(text string) {
			re := regexp.MustCompile(`\d+`)
			rateStr := re.FindString(text)

			if rateStr == "" {
				return
			}

			rating, err := strconv.Atoi(rateStr)
			if err != nil {
				return
			}

			if rating >= 1 && rating <= 5 {
				log.Printf("rating: %d", rating)
				review.Rating = rating
				review.ValidRating = true
			}
		},
	}

	reviewInteraction := conversation.Interaction{
		Statement:      `Hello again! We noticed you've recently received your iPhone 13. We'd love to hear about your experience. Can you spare a few minutes to share your thoughts?`,
		NLPResolver:    conversation.NLPResolverLexicon,
		PositiveAnswer: rateInteraction,
		NegativeAnswer: &conversation.Interaction{
			Statement: `Okay, no problem, can we reach out to you at a later time?`,
			PositiveAnswer: &conversation.Interaction{
				Statement: `Thank you!`,
			},
			NegativeAnswer: &conversation.Interaction{
				Statement: `Okay, no problem, have a great day!`,
			},
		},
		AmbivalentAnswer: &conversation.Interaction{
			Statement:      `In one word, how would you define the iPhone 13?`,
			PositiveAnswer: rateInteraction,
		},
	}

	return &ReviewConversation{
		Conversation: conversation.Conversation{
			Interaction: reviewInteraction,
		},
		Review: review,
	}
}

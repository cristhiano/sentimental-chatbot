package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cristhiano/reviews-chatbot/review"
)

func prompt(msg string) string {
	var input string
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, msg+"\n")
		input, _ = r.ReadString('\n')
		if input != "" {
			break
		}
	}

	return strings.TrimSpace(input)
}

func main() {
	conversation := review.NewReviewConversation("iPhone 13", "1234")
	human, robot := conversation.StartReview()

	for statement := range robot {
		human <- prompt(statement)
	}

	fmt.Println(conversation.Review)
}

package main

import (
	"fmt"

	"github.com/cristhiano/sentimental-chatbot/scripts"
	"github.com/cristhiano/sentimental-chatbot/transport"
)

func main() {
	c := scripts.NewReviewConversation("iPhone 13", "1234")
	cli := new(transport.CLI)

	// blocks until the conversation is over
	c.Start(cli)

	fmt.Println(c.Review)
}

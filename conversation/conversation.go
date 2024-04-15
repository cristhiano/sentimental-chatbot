package conversation

import "github.com/cristhiano/sentimental-chatbot/transport"

type Conversation struct {
	Interaction Interaction
}

func (c *Conversation) Start(t transport.Transport) {
	human := make(chan string)
	robot := make(chan string)

	go c.output(&c.Interaction, human, robot)

	c.input(t, human, robot)
}

func (c *Conversation) input(t transport.Transport, human, robot chan string) {
	for statement := range robot {
		human <- t.Input(statement)
	}
}

func (c *Conversation) output(interaction *Interaction, human, robot chan string) {
	robot <- interaction.Statement

	for answer := range human {
		interaction = interaction.Resolve(answer)

		if interaction == nil {
			close(robot)
			break
		}

		robot <- interaction.Statement
	}
}

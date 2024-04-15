package conversation

type Conversation struct {
	Interaction Interaction
}

func (c *Conversation) Start(human chan string) chan string {
	robot := make(chan string)

	go c.perform(human, robot)

	return robot
}

func (c *Conversation) perform(human, robot chan string) {
	interaction := &c.Interaction
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

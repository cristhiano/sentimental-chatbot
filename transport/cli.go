package transport

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct{}

func (c *CLI) Input(msg string) string {
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

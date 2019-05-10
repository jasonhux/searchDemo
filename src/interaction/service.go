package interaction

import (
	"bufio"
	"fmt"
	"strings"
)

type Service interface {
	GetUserInput() (isQuitCommand bool, input string)
}

type service struct {
	Scanner *bufio.Scanner
}

type Scanner interface {
	Scan() bool
	Text() string
}

func NewService(scanner *bufio.Scanner) Service {
	return &service{scanner}
}

func (s *service) GetUserInput() (isQuitCommand bool, input string) {
	fmt.Print("> ")
	s.Scanner.Scan()
	input = s.Scanner.Text()
	if strings.ToLower(input) == "quit" {
		isQuitCommand = true
	}
	return
}

package interaction

import (
	"bufio"
	"fmt"
)

type Service interface {
	GetUserInput() (isQuitCommand bool, input string)
}
type service struct {
	Scanner *bufio.Scanner
}

func NewService(scanner *bufio.Scanner) Service {
	return &service{scanner}
}

func (s *service) GetUserInput() (isQuitCommand bool, input string) {
	fmt.Print("> ")
	s.Scanner.Scan()
	input = s.Scanner.Text()
	if input == "quit" {
		isQuitCommand = true
	}
	return
}

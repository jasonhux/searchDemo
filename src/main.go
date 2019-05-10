package main

import (
	"bufio"
	"fmt"
	"os"
	"searchDemo/src/data"
	"searchDemo/src/interaction"
	"searchDemo/src/search"
)

func main() {
	dataService := data.NewService()
	interactionService := interaction.NewService(bufio.NewScanner(os.Stdin))

	s := search.NewService(dataService, interactionService)
	for {
		err := s.SetStructMap()
		if err != nil {
			fmt.Println(err)
			return
		}
		isQuit, err := s.Search()
		if isQuit {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		isNewSearchRequired := s.RequestNewSearch()
		if isNewSearchRequired {
			continue
		}
		break

	}

}

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
	err := s.SetStructMap()
	if err != nil {
		fmt.Println(err)
		//stop here
		s.RequestNewSearch()
		return
	}
	for {
		results, isQuit, err := s.Search()
		if isQuit {
			break
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(results)
		}

		isNewSearchRequired := s.RequestNewSearch()
		if isNewSearchRequired {
			continue
		}
		break
	}

}

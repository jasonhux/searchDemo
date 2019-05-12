package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"searchDemo/src/data"
	"searchDemo/src/interaction"
	"searchDemo/src/search"
)

func main() {
	dataService := data.NewService(data.NewSerializer())
	interactionService := interaction.NewService(bufio.NewScanner(os.Stdin))

	s := search.NewService(dataService, interactionService)
	err := s.SetStructMap()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to set the struct map, press any key to exit the application")
		interactionService.GetUserInput()
		return
	}
	for {
		results, isQuit, err := s.StartSearch()
		if isQuit {
			break
		}
		if err != nil {
			fmt.Println(err)
		} else {
			printOutput(results)
		}

		isNewSearchRequired := s.RequestNewSearch()
		if isNewSearchRequired {
			continue
		}
		break
	}
}

func printOutput(v interface{}) {
	resultsBytes, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
	}
	resultsJSONString := string(resultsBytes)
	fmt.Println(resultsJSONString)
}

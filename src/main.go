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

	//Load the struct map into search service before user gets prompts for searches. If load fails, inform user and exit the application
	err := s.SetStructMap()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to set the struct map, press any key to exit the application")
		interactionService.GetUserInput()
		return
	}

	//Loop the StartSearch func so the application will continue to run (either successful or failed search) unless user select to quit
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

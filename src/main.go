package main

import (
	"bufio"
	"fmt"
	"os"
	"searchDemo/src/data"
	"searchDemo/src/search"
	"time"
)

func main() {
	start := time.Now()

	dataService := data.NewService()
	s := search.NewService(dataService)
	err := s.SetStructMap()
	if err != nil {
		fmt.Println(err)
	}
	colapse := time.Now().Sub(start)
	fmt.Println("Load consumed:", colapse)

	scanner := bufio.NewScanner(os.Stdin)

	displayMenu()
	fmt.Println("Select 1) Tickets or 2) Users or 3) Organizations")
	fmt.Print("> ")
	scanner.Scan()
	searchStructParam := scanner.Text()
	fieldMap, err := s.SetSearchStruct(searchStructParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Available search field")
	for k := range fieldMap {
		fmt.Println(k)
	}
	fmt.Println("Enter search field")
	fmt.Print("> ")
	scanner.Scan()
	searchFieldParam := scanner.Text()
	err = s.SetSearchFieldValue(searchFieldParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter search value")
	fmt.Print("> ")
	scanner.Scan()
	searchValueParam := scanner.Text()
	start = time.Now()
	results, err := s.Search(searchValueParam)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(results)
	}
	colapse = time.Now().Sub(start)
	fmt.Println("Search and print time consumed:", colapse)
	scanner.Scan()

}

func displayMenu() {
	fmt.Println("Select search options:")
	fmt.Println("*Press 1 to start search")
	fmt.Println("*Press 2 to view a list of searchable fields")
	fmt.Println("*Type 'quit' to exit")
	fmt.Print("> ")
}

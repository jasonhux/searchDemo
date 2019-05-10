package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"searchDemo/src/search"
	"time"
)

func main() {
	//Load files and unmarshal
	// loadData(&tickets)
	start := time.Now()
	tickets, users, organizations, _ := loadFile()
	structMap := search.PrepareStructMap(tickets, users, organizations)
	colapse := time.Now().Sub(start)
	fmt.Println("Load consumed:", colapse)
	s := search.NewService(structMap)

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
	// resultsList, err := s.Search(searchValueParam)
	_, err = s.Search(searchValueParam)
	if err != nil {
		fmt.Println(err)
		return
	}

	colapse = time.Now().Sub(start)
	fmt.Println("Search and print time consumed:", colapse)
	scanner.Scan()

}

func loadFile() (tickets []*search.Ticket, users []*search.User, organizations []*search.Organization, err error) {
	data, _ := ioutil.ReadFile("tickets.json")
	json.Unmarshal(data, &tickets)
	data, _ = ioutil.ReadFile("users.json")
	json.Unmarshal(data, &users)
	data, _ = ioutil.ReadFile("organizations.json")
	json.Unmarshal(data, &organizations)
	//add error handling
	return
}

func displayMenu() {
	fmt.Println("Select search options:")
	fmt.Println("*Press 1 to start search")
	fmt.Println("*Press 2 to view a list of searchable fields")
	fmt.Println("*Type 'quit' to exit")
	fmt.Print("> ")
}

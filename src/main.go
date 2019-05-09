package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"searchDemo/src/search"
	"time"
)

func main() {
	//Load files and unmarshal
	// loadData(&tickets)
	start := time.Now()
	tickets := loadFile()
	fieldList := search.PrepareTickets(tickets)

	resultsList, err := search.SearchTicket("Status", "hold", fieldList)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, result := range resultsList {
			ticket := result.(*search.Ticket)
			fmt.Println(ticket.ID)
		}
		// fmt.Println(resultsList)
	}
	colapse := time.Now().Sub(start)
	fmt.Println(colapse)

	//buf := bufio.NewReader(os.Stdin)
	// scanner := bufio.NewScanner(os.Stdin)
	// for {
	// 	displayMenu()
	// 	loadFile(tickets)
	// 	//selectSearchOptions(scanner)

	// 	// sentence, err := buf.ReadString('\n')
	// 	// if err != nil {
	// 	// 	fmt.
	// 	// 		Println(err)
	// 	// } else {
	// 	// 	if strings.Contains(sentence, "exit") {
	// 	// 		break
	// 	// 	}
	// 	// 	fmt.Print(sentence)
	// 	// }
	// }
}

func loadFile() []*search.Ticket {
	tickets := []*search.Ticket{}
	data, _ := ioutil.ReadFile("tickets.json")
	json.Unmarshal(data, &tickets)
	return tickets
}

func loadData(tickets *[]search.Ticket) {
	fmt.Println("Loading Data...")
}

func displayMenu() {
	fmt.Println("Select search options:")
	fmt.Println("*Press 1 to start search")
	fmt.Println("*Press 2 to view a list of searchable fields")
	fmt.Println("*Type 'quit' to exit")
	fmt.Print("> ")
}

func selectSearchOptions(buf *bufio.Scanner) {
	// input, _ := buf.ReadBytes('\n')

	//have a mapping table instead of case;
	//check whether it's a number,if yes, whether it's out of range, or otherwise throw error
	buf.Scan()
	switch buf.Text() {
	case "1":
		fmt.Println("You selected user")
		break
	case "2":
		fmt.Println("You selected 2")
		break
	default:
		fmt.Println("default")
	}
}

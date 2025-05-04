package main

import "fmt"

func main() {
	meeting := "meeting on Tuesday"
	appoint := "doctor on Friday"
	party := "Leah's birthday on Saturday"

	var taskItems = []string{meeting, appoint, party}
	fmt.Println("---List of my to-do ---")

	for index, tasks := range taskItems {
		fmt.Printf("%d. %s\n", index+1, tasks)

	}
}

package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("---List of my to-do ---")

	http.HandleFunc("/hello-go", helloUser)

	http.ListenAndServe(":8080", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "hello user. Welcome to our to do list App"
	fmt.Fprintln(writer, greeting)

}

func printTasks(taskItems []string) { //passando parametros e especificando o tipo {
	for index, tasks := range taskItems {
		fmt.Printf("%d. %s\n", index+1, tasks) // placeholders na string"" e na sequencia os values correspondentes...
	}
}

func addTask(taskItems []string, newTask string) {
	var updatedTask = append(taskItems, newTask) // append juntar
	printTasks(updatedTask)
}

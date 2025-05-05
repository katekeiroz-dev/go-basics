package main

import (
	"fmt"
	"net/http"
)

var meeting = "meeting on Tuesday"
var appoint = "doctor on Friday"
var party = "Leah's birthday on Saturday"

var taskItems = []string{meeting, appoint, party} //chama-se slice, pois nao eh limitado os values dentro. se por acaso eu limito ex [2] seria uma array

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

}

func helloUser(writer http.ResponseWriter, request *http.Request) { //endpoints and function handlers
	var greeting = "hello user. Welcome to our to do list App"
	fmt.Fprintln(writer, greeting)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

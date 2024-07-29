package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	command := flag.Arg(0)

	todos, err := Read()
	if err != nil {
		fmt.Printf("Failed to parse file")
		fmt.Printf("%s", err.Error())
		os.Exit(1)
	}

	switch command {
	case "list":
		listCmd(&todos)
	case "add":
		addCmd(&todos)
	case "clear":
		Clear()
	case "toggle":
		toggleCmd(&todos)
	case "update":
		updateCmd(&todos)
	default:
		defaultCmd()
	}
}

func defaultCmd() {
	fmt.Println(`
Godo v0.1
Usage: godo [command]
			
available commands:
  add [title]         Adds a new todo item
  clear		        Clears all todo items
  list                Lists all current todo items
  toggle [id]         Toggles a todo item as done or not done
  update [id] [title] Updates the title of a todo item`)
}

func listCmd(todos *Todos) {
	if len(todos.todos) == 0 {
		fmt.Println("No todos found")
		return
	}

	for _, todo := range todos.todos {
		fmt.Println(todo.String())
	}
}

func addCmd(todos *Todos) {
	title := flag.Arg(1)
	if title == "" {
		fmt.Println("Title is required")
		os.Exit(1)
	}
	todos.Add(title)
	fmt.Println("Todo added")
}

func toggleCmd(todos *Todos) {
	id := flag.Arg(1)
	if id == "" {
		fmt.Println("Id is required")
		os.Exit(1)
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Id must be a number")
		os.Exit(1)
	}

	err = todos.Toggle(intId)
	if err != nil {
		fmt.Println("Todo not found")
		os.Exit(1)
	}
}

func updateCmd(todos *Todos) {
	id := flag.Arg(1)
	if id == "" {
		fmt.Println("Id is required")
		os.Exit(1)
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Id must be a number")
		os.Exit(1)
	}

	title := flag.Arg(2)
	if title == "" {
		fmt.Println("Title is required")
		os.Exit(1)
	}

	err = todos.Update(intId, title)
	if err != nil {
		fmt.Println("Todo not found")
		os.Exit(1)
	}
}

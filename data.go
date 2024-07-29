package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Clear() {
	jsonPath := getPath()
	os.WriteFile(jsonPath, []byte("[]"), 0644)
}

func Read() (Todos, error) {
	var todos []Todo

	jsonPath := getPath()

	_, err := os.Stat(jsonPath)
	if err != nil {
		os.WriteFile(jsonPath, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return Todos{}, err
	}

	err = json.Unmarshal(data, &todos)
	if err != nil {
		return Todos{}, err
	}

	result := Todos{
		todos: todos,
	}
	return result, nil
}

func Write(todos Todos) {
	jsonPath := getPath()
	data, err := json.Marshal(todos.todos)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(jsonPath, data, 0644)
	if err != nil {
		panic(err)
	}
}

func getPath() string {
	cur, _ := os.Executable()
	dir := filepath.Dir(cur)
	return filepath.Join(dir, "godo.json")
}

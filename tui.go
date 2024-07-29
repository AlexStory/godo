package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	todos  Todos
	cursor int
}

func initialModel(todos Todos) model {
	return model{
		todos:  todos,
		cursor: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(m.todos.todos)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case " ":
			m.todos.Toggle(m.todos.todos[m.cursor].Id)
			updated, err := Read()
			if err != nil {
				m.todos = updated
				return m, tea.Quit
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Todos:\nCount:"
	s += fmt.Sprintf("%d\n", len(m.todos.todos))
	for i, todo := range m.todos.todos {
		if i == m.cursor {
			s += fmt.Sprintf("> %s\n", todo.String())
		} else {
			s += fmt.Sprintf("  %s\n", todo.String())
		}
	}
	s += "j/k to move, space to toggle, q to quit"
	return s
}

func runTui(todos Todos) {
	p := tea.NewProgram(initialModel(todos))
	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

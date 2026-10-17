package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"tdl/todolist/tasklist"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := make(tasklist.TaskList)

	fmt.Println("Welcome to ToDoList. Type help to get list of commands. Type exit to exit")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if input == "exit" {
			break
		}

		if err := ParseInput(input, &tasks); err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("input error:", err)
	}
}

func ParseInput(input string, tl *tasklist.TaskList) error {
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return nil
	}

	switch parts[0] {
	case "help":
		PrintHelp()

	case "list":
		l := tl.List()
		tasklist.PrintList(l)

	case "add":
		if len(parts) < 3 {
			return errors.New("usage: add <name> <text>")
		}

		name := parts[1]
		text := strings.Join(parts[2:], " ")

		if err := tl.Add(name, text); err != nil {
			return err
		}

		fmt.Println("task added")

	case "done":
		if len(parts) < 2 {
			return errors.New("usage: done <name>")
		}

		name := parts[1]

		if err := tl.CompleteTask(name); err != nil {
			return err
		}

		fmt.Println("task done")

	case "del":
		if len(parts) < 2 {
			return errors.New("usage: del <name>")
		}

		name := parts[1]

		if err := tl.Delete(name); err != nil {
			return err
		}

		fmt.Println("task deleted")

	default:
		fmt.Println("type help to get list of available commands")
	}

	return nil
}

func PrintHelp() {
	commands := `help                   - Display the list of available commands and their syntax
add <name> <text>      - Add a new task
list                   - Display all tasks
del <name>             - Delete a task by its title
done <name>            - Mark a task as completed
events                 - Display the event log
exit                   - Exit the program`

	fmt.Println(commands)
}

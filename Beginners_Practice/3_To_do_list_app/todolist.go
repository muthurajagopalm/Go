package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func displayMenu() {

	fmt.Println("\n To-Do-List Menu")
	fmt.Println("1. Add Task")
	fmt.Println("2. View Task")
	fmt.Println("3. Remove Task")
	fmt.Println("4. Exit")
}

func addTask(file *os.File) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the task:")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	_, err := file.WriteString(task + "\n")

	if err != nil {
		fmt.Println("Error writing task to a file:", err)
		return
	}
	fmt.Println("Task added successfully!")
}

func main() {

	file, err := os.OpenFile("todo.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer file.Close()
	displayMenu()
	addTask(file)

}

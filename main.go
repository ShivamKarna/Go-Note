package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Println("Add a new note ? (y/n) : ")

		userInput, _ := inputReader.ReadString('\n')

		userInput = strings.TrimSpace(userInput)

		if userInput == "n" {
			fmt.Println("Ok GoodBye!")
			break
		} else {
			reader := bufio.NewReader(os.Stdin)

			fmt.Println("Enter Your note : ")

			noteContent, err := reader.ReadString('\n')

			noteContent = strings.TrimSpace(noteContent)

			if err != nil {
				fmt.Println("Error occured while reading input")
			}

			// Open or create file and write to it
			file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644) // 0644 for the permissoins in linux to give write and read permissions
			if err != nil {
				fmt.Println("Error occured while opening file")
			}

			defer file.Close()

			// actually write to file now
			if _, err := file.WriteString(noteContent + "\n"); err != nil {
				fmt.Println("Error occured while writing to file")
			} else {
				fmt.Println("Note content written to file")
			}
		}

	}
	// this is the main closing braces
}


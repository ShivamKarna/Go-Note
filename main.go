package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("\n--- NOTE DASHBOARD ---")
		fmt.Println("1. Add a New Note")
		fmt.Println("2. View All Notes") // chore for later!
		fmt.Println("3. Delete Notes File")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choiceReader := bufio.NewReader(os.Stdin)
		choice, _ := choiceReader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Enter Your note : ")
			inputReader := bufio.NewReader(os.Stdin)

			noteContent, err := inputReader.ReadString('\n')

			noteContent = strings.TrimSpace(noteContent)

			if err != nil {
				fmt.Println("Error occured while reading input")
			}

			// Open or create file and write to it
			file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644) // 0644 for the permissoins in linux to give write and read permissions
			if err != nil {
				fmt.Println("Error occured while opening file")
			}

			// actually write to file now
			if _, err := file.WriteString(noteContent + "\n"); err != nil {
				fmt.Println("Error occured while writing to file")
			} else {
				fmt.Println("Note content written to file")
			}
			file.Close()

		case "2":
			fmt.Println("Viewing all Notes...") // TODO: Left to do

		case "3":
			fmt.Println("Are you sure you want to delete the notes files ? (y/n) :")
			deleteReader := bufio.NewReader(os.Stdin)
			deleteChoice, _ := deleteReader.ReadString('\n')

			deleteChoice = strings.TrimSpace(deleteChoice)

			if deleteChoice == "y" {
				err := os.Remove("notes.txt")
				if err != nil {
					fmt.Println("Error while deleting file, and the error is : ", err)
				} else {
					fmt.Println("Notes file delted successfully")
				}
			} else {
				fmt.Println("Delete cancelled")
			}
		case "4":
			fmt.Println("Exiting Program")
			return
		default:
			fmt.Println("Invalid choice, please pick 1-4.")
		}

	}
	// this is the main closing braces
}

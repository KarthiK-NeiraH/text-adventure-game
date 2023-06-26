package main

import (
	"bufio"
	"fmt"
	"os"
)

type Room struct {
	Name        string
	Description string
	Exits       map[string]*Room
}

func main() {
	// Create rooms
	entrance := &Room{
		Name:        "Entrance",
		Description: "You are standing at the entrance of a dark cave.",
		Exits:       make(map[string]*Room),
	}

	hallway := &Room{
		Name:        "Hallway",
		Description: "You are in a long hallway with torches on the walls.",
		Exits:       make(map[string]*Room),
	}

	treasureRoom := &Room{
		Name:        "Treasure Room",
		Description: "You have found the hidden treasure room!",
		Exits:       make(map[string]*Room),
	}

	// Connect rooms
	entrance.Exits["north"] = hallway
	hallway.Exits["south"] = entrance
	hallway.Exits["east"] = treasureRoom
	treasureRoom.Exits["west"] = hallway

	// Start the game
	currentRoom := entrance

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Text Adventure Game!")
	fmt.Println("You find yourself in a dark cave. Your goal is to find the hidden treasure.")
	fmt.Println("Enter 'quit' at any time to exit the game.")

	for {
		fmt.Println()
		fmt.Println(currentRoom.Name)
		fmt.Println(currentRoom.Description)

		if currentRoom == treasureRoom {
			fmt.Println("Congratulations! You found the treasure. You win!")
			break
		}

		fmt.Print("Choose an exit: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "quit" {
			fmt.Println("Thank you for playing. Goodbye!")
			break
		}

		nextRoom, found := currentRoom.Exits[input]
		if !found {
			fmt.Println("You can't go that way. Try again.")
		} else {
			currentRoom = nextRoom
		}
	}
}

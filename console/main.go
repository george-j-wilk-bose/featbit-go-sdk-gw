package main

import (
	"fmt"

	"github.com/featbit/featbit-go-sdk"
)

func main() {
	envSecret := "9UX39K2_2EOqodqzqvdKZQL8ntSptm7EaTln2Ccmy4og"
	streamingUrl := "ws://localhost:5100"
	eventUrl := "http://localhost:5100"
	client, err := featbit.NewFBClient(envSecret, streamingUrl, eventUrl)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	defer func() {
		if client != nil {
			_ = client.Close()
		}
	}()

	if !client.IsInitialized() {
		fmt.Println("SDK initialization failed")
		return
	}

	for {
		fmt.Print("Press 'g' to get metadata, 'q' to quit: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		switch input {
		case "q":
			fmt.Println("Quitting...")
			return
		case "g":
			metadata, err := client.GetAllFlagMetadata()
			if err != nil {
				fmt.Printf("Error getting metadata: %v\n", err)
			} else {
				for _, meta := range metadata {
					fmt.Printf("flag metadata: %+v\n", meta)
				}
			}
		default:
			fmt.Println("Invalid input. Use 'g' for metadata or 'q' to quit.")
		}
	}
}

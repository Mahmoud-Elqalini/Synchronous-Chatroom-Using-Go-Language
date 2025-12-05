package main

import (
	"bufio"   // For buffered I/O (reading/writing text)
	"fmt"     // For printing messages to the console
	"log"     // For logging errors
	"net"     // For network connections (TCP)
	"os"      // For accessing OS features like stdin
	"strings" // For string manipulation (e.g., trimming newline)
)

func main() {
	// -------------------------------
	// 1. Connect to the server
	// -------------------------------
	// Here, we connect to a TCP server running on localhost at port 1234.
	conn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		log.Fatalf("Error connecting to server: %v", err)
	}
	// Make sure to close the connection when the program exits
	defer conn.Close()

	// Create a reader to read input from the user (keyboard)
	reader := bufio.NewReader(os.Stdin)

	// -------------------------------
	// 2. Ask the user for their ID
	// -------------------------------
	fmt.Print("Enter your user ID: ")
	id, _ := reader.ReadString('\n')         // Read input until Enter key
	id = strings.TrimSpace(id)               // Remove trailing newline or spaces

	// -------------------------------
	// 3. Send the ID to the server
	// -------------------------------
	// This lets the server know who just connected
	fmt.Fprintf(conn, "%s\n", id)

	// -------------------------------
	// 4. Welcome message
	// -------------------------------
	fmt.Println("Welcome to the chat")

	// -------------------------------
	// 5. Start reading and writing concurrently
	// -------------------------------
	// Use a goroutine to read messages from the server
	go readFromServer(conn)

	// Main goroutine will handle sending messages to the server
	writeToServer(conn)
}

// -------------------------------
// Function: readFromServer
// -------------------------------
// Listens for messages from the server and prints them to the console
func readFromServer(conn net.Conn) {
	scanner := bufio.NewScanner(conn) // Scanner reads input line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Print each message received from server
	}
	// Handle errors or disconnection
	if err := scanner.Err(); err != nil {
		log.Printf("Disconnected from server: %v", err)
		os.Exit(0)
	}
}

// -------------------------------
// Function: writeToServer
// -------------------------------
// Reads messages from the user and sends them to the server
func writeToServer(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)    // Read user input line by line
	writer := bufio.NewWriter(conn)          // Writer to send text to the server

	for scanner.Scan() {
		text := scanner.Text()               // Get the user's message
		
		// Send the message to the server
		_, err := writer.WriteString(text + "\n")
		if err != nil {
			log.Println("Error sending message:", err)
			return
		}
		writer.Flush()                        // Ensure the message is actually sent

		// Exit condition
		if text == "exit" {
			fmt.Println("Goodbye")
			return
		}
	}
}

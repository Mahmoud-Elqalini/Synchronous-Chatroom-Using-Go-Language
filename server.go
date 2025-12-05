package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

// -------------------------------
// Message structure
// -------------------------------
// This structure holds a message to be broadcasted.
// senderConn is used to prevent sending the message back to the same client (No Self-Echo)
type Message struct {
	senderConn net.Conn
	content    string
}

// -------------------------------
// Global variables
// -------------------------------
var (
	clients   = make(map[net.Conn]string) // Map to store client connections and their IDs
	broadcast = make(chan Message)        // Central channel to receive and distribute messages
	mu        sync.Mutex                  // Mutex to protect the clients map (Thread safety)
)

func main() {
	// -------------------------------
	// 1. Start server on port 1234
	// -------------------------------
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 5555")

	// -------------------------------
	// 2. Start the broadcasting goroutine
	// -------------------------------
	go handleMessages()

	// -------------------------------
	// 3. Accept new client connections
	// -------------------------------
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle each new client in a separate goroutine
		go handleClient(conn)
	}
}

// -------------------------------
// Function: handleMessages
// -------------------------------
// This goroutine continuously reads messages from the broadcast channel
// and sends them to all connected clients (except the sender)
func handleMessages() {
	for msg := range broadcast {
		mu.Lock() // Lock to safely access clients map
		for conn := range clients {
			if conn != msg.senderConn { // Prevent sending back to the sender
				fmt.Fprintln(conn, msg.content) // Send the message
			}
		}
		mu.Unlock()
	}
}

// -------------------------------
// Function: handleClient
// -------------------------------
// Handles communication with a single client
func handleClient(conn net.Conn) {
	defer conn.Close() // Ensure connection closes when function exits

	reader := bufio.NewReader(conn)

	// -------------------------------
	// 1. Read User ID as the first input
	// -------------------------------
	idStr, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	id := strings.TrimSpace(idStr)

	// -------------------------------
	// 2. Register the client
	// -------------------------------
	mu.Lock()
	clients[conn] = id
	mu.Unlock()

	// Notify all clients that a new user joined
	joinMsg := fmt.Sprintf("> User %s joined", id)
	broadcast <- Message{
		senderConn: conn, // prevent sending join message back to the same user
		content:    joinMsg,
	}

	// -------------------------------
	// 3. Receive messages from this client
	// -------------------------------
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break // Client disconnected
		}
		msg = strings.TrimSpace(msg)

		// Broadcast non-empty messages
		if msg != "" {
			formattedMsg := fmt.Sprintf("[%s]: %s", id, msg)
			broadcast <- Message{
				senderConn: conn,
				content:    formattedMsg,
			}
		}
	}

	// -------------------------------
	// 4. Handle client disconnect
	// -------------------------------
	mu.Lock()
	delete(clients, conn) // Remove client from map
	mu.Unlock()

	leaveMsg := fmt.Sprintf("User %s left the chat", id)
	broadcast <- Message{
		senderConn: nil, // nil because we don't care about Self-Echo for leaving messages
		content:    leaveMsg,
	}

	log.Printf("User %s disconnected.", id)
}

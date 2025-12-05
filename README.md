# Simple Chatroom (Go RPC)

## 📖 Description

This project implements a **multi-client chatroom** using **Go and TCP sockets**.  

It allows multiple clients to connect to a central server, exchange messages in real time, and automatically share join/leave notifications.  

The chatroom works as a **text-based console application**.

---

## Features

### 🧠 Server
- Handles multiple clients over TCP concurrently.
- Stores connected clients and manages join/leave notifications.
- Broadcasts messages to all connected clients except the sender.
- Thread-safe handling of clients using `sync.Mutex`.

### 💬 Client
- Connects to the server via TCP.
- Sends messages and receives messages from other clients in real time.
- Displays notifications when users join or leave the chat.
- Supports graceful exit using the `exit` command.
- Shows messages like:
```
> User Mahmoud joined
[Mahmoud]: Hello everyone!
User Mahmoud left the chat
```

---


## 🗂️ Folder Structure
```
project/
│
├── server.go        # Chat server logic (handles Join, Leave, SendMessage)
├── client.go        # Chat client logic (connects and communicates via RPC)
├── commons/
│   └── args.go      # Shared data structures (MessageArgs)
├── go.mod           # Go module file
├── .gitignore       # Git ignore configuration
└── README.md        # Project documentation
```

---

## ⚙️ How to Run

### 1️⃣ Start the Server
```bash
go run server.go
```

### 2️⃣ Start the Client (in a new terminal)
```bash
go run client.go
```

### 3️⃣ Interact with the Chat
- Type your **user ID** when prompted and press **Enter**.
- Type your message and press **Enter** to send.
- Type **exit** to leave the chatroom.
- When a user joins or leaves, everyone sees a server message automatically.

---

## 🧩 Example Output

**Client 1 (Mahmoud):**
```
Enter your user ID: Mahmoud
Welcome to the chat!
Hello everyone!
```
**Other clients see:**
```
> User Mahmoud joined
[Mahmoud]: Hello everyone!
```
**When Mahmoud leaves:**
```
User Mahmoud left the chat
```

---


## Technical Notes
- Uses **net** package for TCP connections.
- Each client is handled in a separate goroutine.
- Broadcast channel is used to distribute messages to all clients.
- Thread-safe access to shared client list with `sync.Mutex`.
- Prevents **self-echo** (clients do not see their own messages twice).

---

## 🎥 Demo Video
🔗 Watch the running application demo here: *([video link here](https://drive.google.com/file/d/1J20QXFc4HrKfLxk7Yg01Jix3WCX6hF42/view?usp=drive_link))*

---

## 👨‍💻 Author
**Mahmoud Hamdi**  
Computer Engineering Student @ Tanta University  
Passionate about software engineering, data science, and AI.

---


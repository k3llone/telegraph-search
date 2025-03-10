A project for finding links on the Telegraph website by iterating through numbers in the URL. The project includes a web interface written in HTML/CSS/JS and a backend written in Go using the Gin and Gorilla/WebSocket libraries.

## Description

The project allows you to find public pages on Telegraph by iterating through possible number combinations in the URL. The web interface provides a convenient way to interact with the tool, and WebSocket ensures dynamic result updates.

## Features

- **Link Iteration**: Automatic iteration through numbers in the URL to find existing pages.
- **Web Interface**: A user-friendly interface for managing the search process and displaying results.
- **WebSocket**: Real-time data updates in the web interface.
- **Go**: The backend is written in Go using Gin for routing and Gorilla/WebSocket for WebSocket communication.

## Installation and Setup

### Prerequisites

- Installed Go (version 1.16 or higher)

### Cloning the Repository

```bash
git clone https://github.com/k3llone/telegraph-search.git
cd telegraph-search
```

## Usage

Install dependencies:
```bash
go mod download
```
Run the server: 
```bash
go run ./src
```

The server will start at http://localhost:8080

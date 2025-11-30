package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	dirPath := "."

	// Check if the directory exists
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		fmt.Println("Directory '%s', not found.\n", dirPath)
		return
	}

	// Create a file server handler to serve the directory's contents
	fileServer := http.FileServer(http.Dir(dirPath))

	// Create a new HTTP server and handle requests
	http.Handle("/", fileServer)

	// Start the server on port 8080
	port := 8080
	fmt.Printf("Server started at http://localhost:%d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting server: %s\n", err)
	}
}

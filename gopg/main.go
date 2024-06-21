package main

import (
	"log"
	"os"
)

func main() {
	// create a new logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Log a message
	logger.Println("This is a log message.")

	//Log with formatting
	logger.Printf("Logging with formatted message: %s\n", "Hello World!")
}

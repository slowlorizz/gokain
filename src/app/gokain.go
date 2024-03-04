package main

import (
	"fmt"
	"gokain/logs"
)

func main() {
	args := Args{}
	args.Load()

	fmt.Println("Starting Log-Handler...")
	logs.Handler.Start(args.verbose) // starts the async Log-Handler

	logs.Debug("Hello World", "a", "b", "c", 4)

	logs.Handler.Join() // Prints the last received Log-Message before program-end
}

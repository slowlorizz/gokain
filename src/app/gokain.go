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

	logs.Handler.Join() // Prints the last received Log-Message before program-end
}

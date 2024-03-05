package main

import (
	"fmt"
	"gokain/types"
	utils "gokain/utils"
	"time"

	feat "gokain/feat"

	logs "github.com/slowlorizz/gokain-logs"
)

//  go mod edit -replace gokain/logs=../lib/logs

// 91c9aeb857d46cc8c975c58b0330fcfcd0e01ca3

func main() {
	args := utils.Args{}
	args.Load()
	logs.Handler.Start(args.Verbose) // starts the async Log-Handler

	logs.Info("Log-Handler started!")

	if len(args.Hash) < 1 {
		logs.Fatal("No Hash in Args [-]")
		return
	} else {
		logs.Debug("Hash in args [+]", "hash", args.Hash)
	}

	foundCH := make(chan types.ResultPair)
	stopCH := make(chan bool)

	var result types.ResultPair

	seeds := DistributeSeeds(args.Threads)

	for i := 0; i < args.Threads; i++ {
		go feat.CrackHash(stopCH, foundCH, args.Hash, args.HashType, seeds[i])
	}

join_loop:
	for {
		select {
		case v := <-foundCH:
			result = v
			logs.Info("Found", "plaintext", v.PlainText, "hash", v.Hash)
			break join_loop
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

	stopCH <- true

	logs.Handler.Join() // Prints the last received Log-Message before program-end

	fmt.Printf("\n\n#---------------------------------------------------------------------------------------------------------------------#\n\n")
	fmt.Printf("|  Plain-Text: \033[38;2;0;255;0m%s\033[0m  | Hash: \033[38;2;0;128;255m%s\033[0m  |\n", result.PlainText, result.Hash)
	fmt.Printf("\n\n#---------------------------------------------------------------------------------------------------------------------#\n\n")
	fmt.Println(" ")
}

func DistributeSeeds(threads int) []string {
	charMap := types.GenCharMap(true, true, true, true)
	seeds := make([]string, threads)

	for i := 0; i < len(charMap); i++ {
		seeds[i%threads] += string(charMap[i])
	}

	return seeds
}

package main

import (
	"gokain/types"
	utils "gokain/utils"

	"log"

	ui "github.com/gizak/termui/v3"

	feat "gokain/feat"
	tui "gokain/tui"
)

//  go mod edit -replace gokain/logs=../lib/logs

// 91c9aeb857d46cc8c975c58b0330fcfcd0e01ca3

func main() {
	args := utils.Args{}
	args.Load()

	if len(args.Hash) < 1 {
		log.Fatal("No Hash in Args [-]")
		return
	}

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	foundCH := make(chan types.ResultPair)
	stopCH := make(chan bool)

	//var result types.ResultPair

	seeds := DistributeSeeds(args.Threads)

	height := 3

	thread_components := make([]*tui.ThreadComponent, args.Threads)

	for i := 0; i < args.Threads; i++ {
		thc := tui.New_ThreadComponent(i+1, seeds[i], 1, (height+1)*i)
		thread_components[i] = thc
		go feat.CrackHash(stopCH, foundCH, args.Hash, args.HashType, seeds[i], thc)
	}

	uiEvents := ui.PollEvents()
	//main_loop:
	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			}
		case <-foundCH:
			//result = v
			//close(foundCH)
			stopCH <- true
			close(stopCH)
			for _, tc := range thread_components {
				tc.Render()
			}
			//break main_loop
		default:
			for _, tc := range thread_components {
				tc.Render()
			}
		}
	}
}

func DistributeSeeds(threads int) []string {
	charMap := types.GenCharMap(true, true, true, true)
	seeds := make([]string, threads)

	for i := 0; i < len(charMap); i++ {
		seeds[i%threads] += string(charMap[i])
	}

	return seeds
}

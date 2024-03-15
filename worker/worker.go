package main

import (
	"fmt"

	"github.com/slowlorizz/gokain/worker/src/args"
	"github.com/slowlorizz/gokain/worker/src/thread"
	"github.com/slowlorizz/gokain/worker/src/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/src/ui"

	tui "github.com/gizak/termui/v3"
)

var Result thread.ResultPair

func main() {
	err := args.Load()

	if err != nil {
		err.Raise()
	}

	ui.Init()

	chrs := charset.New(true, true, true, true, false, false, false)

	seeds := make([][]string, args.Threads)

	for i := 0; i < args.Threads; i++ {
		seeds[i] = make([]string, 0)
	}

	for i, c := range chrs.Chars {
		seeds[i%args.Threads] = append(seeds[i%args.Threads], c)
	}

	for i, v := range seeds {
		th := thread.New(i+1, args.Hash, v, args.HashType, *chrs, thread.FOUND_CH, thread.JOIN_CH)
		ui.Components = append(ui.Components, &th.UiC)
		go th.Start()
	}

	ui.Events = tui.PollEvents()
	for {
		select {
		case e := <-ui.Events:
			switch e.ID {
			case "q", "<C-c>":
				thread.StopAll()
				tui.Close()

				if Result.PlainText != "" && Result.Hash != "" {
					fmt.Printf("\nHASH:  %s\nTEXT:  %s\n", Result.Hash, Result.PlainText)
				}

				return
			}
		case res := <-thread.FOUND_CH:
			Result = res
			thread.StopAll()
			ui.Render()
		default:
			ui.Render()
		}
	}
}

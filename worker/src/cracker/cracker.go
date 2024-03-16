package cracker

import (
	"time"

	"github.com/slowlorizz/gokain/worker/src/args"
	"github.com/slowlorizz/gokain/worker/src/thread"
	"github.com/slowlorizz/gokain/worker/src/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/src/ui"

	tui "github.com/gizak/termui/v3"
)

var Result thread.ResultPair
var chrs charset.CharSet = *charset.New(true, true, true, true, false, false, false)
var seeds [][]string
var StartTime time.Time
var RunTime time.Duration
var CtrlCH chan struct{}

func Start() {
	DistributeSeeds()
	for i, v := range seeds {
		th := thread.New(i+1, args.Hash, v, args.HashType, chrs, thread.FOUND_CH, thread.JOIN_CH)
		ui.Components = append(ui.Components, th.UiC)
		go th.Start()
	}

	ui.Clock.Init()

	ui.Events = tui.PollEvents()
	for {
		select {
		case e := <-ui.Events:
			switch e.ID {
			case "q", "<C-c>":
				thread.StopAll()
				close(CtrlCH)
				return
			}
		case res := <-thread.FOUND_CH:
			Result = res
			ui.Clock.Stop = true
			thread.StopAll()
			ui.Render()
		default:
			time.Since(StartTime)
			ui.Render()
		}
	}
}

func DistributeSeeds() {
	seeds = make([][]string, args.Threads)

	for i := 0; i < args.Threads; i++ {
		seeds[i] = make([]string, 0)
	}

	for i, c := range chrs.Chars {
		seeds[i%args.Threads] = append(seeds[i%args.Threads], c)
	}
}

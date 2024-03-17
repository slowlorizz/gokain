package cracker

import (
	"time"

	"github.com/slowlorizz/gokain/worker/src/args"
	"github.com/slowlorizz/gokain/worker/src/job"
	"github.com/slowlorizz/gokain/worker/src/thread"
	"github.com/slowlorizz/gokain/worker/src/ui"

	tui "github.com/gizak/termui/v3"
)

var Result thread.ResultPair
var seeds [][]string
var StartTime time.Time
var RunTime time.Duration
var CtrlCH chan struct{}
var Threads []*thread.Thread = make([]*thread.Thread, 0)
var Job *job.Job

func Start(j *job.Job) {
	/* 16.03.2024
	>B1a% --> 50min 20s (sha256)
	*/
	/*BuildCharset(map[string][]string{
		"standard": {"numbers", "lower_case", "upper_case", "special"},
		"extended": {"lower_case", "upper_case", "special"},
	})*/
	DistributeSeeds()
	StartThreads()
	Handle()
}

func DistributeSeeds() {
	seeds = make([][]string, args.Threads)

	for i := 0; i < args.Threads; i++ {
		seeds[i] = make([]string, 0)
	}

	for i, c := range Job.Chars {
		seeds[i%args.Threads] = append(seeds[i%args.Threads], c)
	}
}

func StartThreads() {
	for i, v := range seeds {
		th := thread.New(i+1, args.Hash, v, args.HashType, Job.Chars, thread.FOUND_CH, thread.JOIN_CH)
		Threads = append(Threads, th)
		ui.Components = append(ui.Components, th.UiC)
		go th.Start()
	}

	ui.Clock.Init()
}

func Handle() {
	ui.Events = tui.PollEvents()
	for {
		select {
		case e := <-ui.Events:
			switch e.ID {
			case "q", "<C-c>":
				thread.StopAll()
				Stop()
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

func Stop() {
	close(CtrlCH)
}

package main

import (
	"log"

	"github.com/slowlorizz/gokain/worker/src/args"
	"github.com/slowlorizz/gokain/worker/src/cracker"
	"github.com/slowlorizz/gokain/worker/src/job"

	tui "github.com/gizak/termui/v3"
)

func main() {
	err := args.Load()

	if err != nil {
		err.Raise()
	}

	if err := tui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	defer tui.Close()

	//------------------

	go cracker.Start(job.New(args.JobPath))

	//------------------
	cracker.Await()
}

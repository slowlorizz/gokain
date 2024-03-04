package gkwrk

import (
	"crypto/sha1"
	"fmt"

	logs "github.com/slowlorizz/gokain-logs"
)

type Worker struct {
	charMap      string // set at runtime
	CipherText   string
	seeds        string
	computeLayer *ComputeLayer      // set at runtime
	outputLayer  *OutputLayer       // set at runtime
	outCh        (chan *ResultPair) // set at runtime
	stopCh       (chan bool)        // set at runtime
	Nums         bool
	Lower        bool
	Upper        bool
	Special      bool
}

func (wrk *Worker) Start(maxRoutines int) {
	logs.Info("Starting Worker")

	wrk.genCharMap()

	wrk.outCh = make(chan *ResultPair)
	wrk.stopCh = make(chan bool)
	wrk.seeds = wrk.charMap

	wrk.computeLayer = &ComputeLayer{
		charMap:     wrk.charMap,
		seeds:       wrk.seeds,
		maxRoutines: maxRoutines,
		outCh:       &wrk.outCh,
		stopCh:      &wrk.stopCh,
		Action:      ComputeAction,
	}

	wrk.outputLayer = &OutputLayer{
		inCh:       &wrk.outCh,
		stopCh:     &wrk.stopCh,
		cipherText: wrk.CipherText,
	}

	wrk.outputLayer.Start()
	wrk.computeLayer.Start()
}

func (wrk *Worker) genCharMap() {
	wrk.charMap = ""

	if wrk.Nums {
		wrk.charMap += "0123456789"
	}

	if wrk.Lower {
		wrk.charMap += "abcdefghijklmnopqrstuvwxyz"
	}

	if wrk.Upper {
		wrk.charMap += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if wrk.Special {
		wrk.charMap += "<>;,:.-_${}[]!^~`'?=()/&%*\"#@|"
	}
}

func (wrk *Worker) Join() {
	for {
		stop := <-wrk.stopCh

		if stop {
			break
		}
	}
}

func ComputeAction(combo *string) *ResultPair {
	res := &ResultPair{PlainText: *combo}

	sha := sha1.New()

	sha.Write([]byte(*combo))

	res.Computed = fmt.Sprintf("%x", sha.Sum(nil))

	return res
}

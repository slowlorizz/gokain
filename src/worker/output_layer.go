package gkwrk

import (
	"fmt"

	logs "github.com/slowlorizz/gokain-logs"
)

type OutputLayer struct {
	cipherText string
	inCh       *(chan *ResultPair)
	stopCh     *(chan bool)
}

func (outp *OutputLayer) Start() {
	go outp.Handle()
}

func (outp *OutputLayer) Handle() {
	for {
		select {
		case <-*outp.stopCh:
			return

		case v := <-*outp.inCh:
			if v.Computed == outp.cipherText {
				logs.Println(fmt.Sprintf("\nFound Value:  \033[38;2;0;255;0m%s  |  %s\033[0m", v.PlainText, v.Computed))
				*outp.stopCh <- true
				return
			} else {
				logs.Println(fmt.Sprintf("%s  |  %s\n", v.PlainText, v.Computed))
			}
		}
	}
}

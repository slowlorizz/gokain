package gkwrk

import (
	"gokain/logs"
)

/* I think the best Configuration for the Allowed Routines-Amount is:
combo-Layer: N --> the least
compute-Layer: 2 * N  --> the most
output-Layer: 2 * N --> Can variy, depens on the task of the output,
						for example:
						if the output is a Rest-Request: wayyy more
						if the output is a console print: probably less
*/

type Worker struct {
	comboLayer   *ComboLayer
	computeLayer *ComputeLayer
	outputLayer  *OutputLayer
	// Pipline mit channels: comboLayer --comboOutCh--> computeLayer --computeOutCh--> outputLayer
	comboOutCh   (chan string)
	computeOutCh (chan string)
}

func (wrk *Worker) Start() {
	logs.Info("Starting Worker")
}

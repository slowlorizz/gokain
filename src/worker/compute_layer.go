package gkwrk

type ComputeLayer struct {
	allowedRoutines uint // How many concurent routines are allowed to run
	inCh            *(chan string)
	outCh           *(chan string)
}

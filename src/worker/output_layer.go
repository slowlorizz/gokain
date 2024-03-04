package gkwrk

type OutputLayer struct {
	allowedRoutines uint // How many concurent routines are allowed to run
	inCh            *(chan string)
}

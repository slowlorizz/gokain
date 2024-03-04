package gkwrk

type ComboLayer struct {
	charMap         string
	allowedRoutines uint // How many routines are allowed to run at the same time
	outCh           *(chan string)
	retCh           (chan bool) // Combinators send true, when finished so ComboLayer knows to start a new Combinator
}

func (cbl *ComboLayer) genCharMap(nums bool, lower bool, upper bool, special bool) {
	cbl.charMap = ""

	if nums {
		cbl.charMap += "0123456789"
	}

	if lower {
		cbl.charMap += "abcdefghijklmnopqrstuvwxyz"
	}

	if upper {
		cbl.charMap += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if special {
		cbl.charMap += "<>;,:.-_${}[]!^~`'?=()/&%*\"#@|"
	}
}

func (cbl *ComboLayer) combinator(seed string, toLen uint) {

}

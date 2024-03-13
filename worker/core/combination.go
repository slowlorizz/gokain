package core

import "strings"

type Combination struct {
	FirstWheel *Wheel
	Seeds      *Wheel
	Chars      []string
}

func New_Combination(chrs string, seeds string) Combination {
	c := Combination{Chars: strings.Split(chrs, "")}
	s := strings.Split(seeds, "")

	c.FirstWheel = New_Wheel(&c.Chars)
	c.Seeds = New_Wheel(&s)

	return c
}

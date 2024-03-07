package core

import "strings"

////////////////////////////////////////////////////////////////////////
// WhlItem

////////////////////////////////////////////////////////////////////////
// Wheel

func New_Wheel(from *[]string) *Wheel {
	whl := Wheel{}
	whl.Chars = from

	var prev *WhlItem
	for _, v := range *whl.Chars {
		itm := &WhlItem{Value: v}

		if whl.First == nil {
			whl.First = itm
			prev = whl.First
		} else {
			prev.Next = itm
			prev = itm
		}
	}

	prev.Next = whl.First
	whl.Current = whl.First

	return &whl
}

func (whl *Wheel) Turn() {
	whl.Current = whl.Current.Next
}

func (whl *Wheel) AtFirst() bool {
	return (whl.Current.Value == whl.First.Value)
}

func (whl *Wheel) AtLast() bool {
	return (whl.Current.Next.Value == whl.First.Value)
}

func (whl *Wheel) Shift() string {
	// Used For Combination, and not Seed

	turnNext := whl.AtLast()
	str := whl.Current.Value
	whl.Turn()

	if turnNext && whl.Next == nil {
		if whl.Seeds.AtLast() {
			whl.Next = New_Wheel(whl.Chars)
			whl.Next.Seeds = whl.Seeds
		}

		whl.Seeds.Turn()
	} else if turnNext {
		str += whl.Next.Shift()
	}

	return str
}

func (whl *Wheel) GetValues() string {
	if whl.Next == nil {
		return whl.Current.Value
	} else {
		return whl.Current.Value + whl.Next.GetValues()
	}
}

////////////////////////////////////////////////////////////////////////
// Combination

func New_Combination(chrs string, seeds string) *Combination {
	c := Combination{Chars: strings.Split(chrs, "")}
	s := strings.Split(seeds, "")

	c.Seed = *New_Wheel(&s)

	c.First = *New_Wheel(&c.Chars)
	c.First.Seeds = &c.Seed

	return &c
}

func (c *Combination) String() string {
	return c.Seed.Current.Value + c.First.GetValues()
}

func (c *Combination) NextCombo() string {
	return c.Seed.Current.Value + c.First.Shift()
}

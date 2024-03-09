package core

type (
	Item struct {
		Value    string
		NextItem *Item
	}

	Wheel struct {
		FirstItem   *Item
		CurrentItem *Item
		NextWheel   *Wheel
	}
)

func New_Wheel(from *[]string) *Wheel {
	whl := &Wheel{}

	var p *Item
	for _, v := range *from {
		itm := &Item{Value: v}

		if whl.FirstItem == nil {
			whl.FirstItem = itm
		} else {
			p.NextItem = itm
		}

		p = itm
	}

	p.NextItem = whl.FirstItem
	whl.CurrentItem = whl.FirstItem

	return whl
}

func (whl *Wheel) AtFirst() bool {
	return (whl.CurrentItem.Value == whl.FirstItem.Value)
}

func (whl *Wheel) AtLast() bool {
	return (whl.CurrentItem.NextItem.Value == whl.FirstItem.Value)
}

func (whl *Wheel) Shift() bool {
	/*
		Goes to next item in Chain
		returns true if it jumped from last to First (did a whole round)
	*/
	whl.CurrentItem = whl.CurrentItem.NextItem
	return whl.AtFirst()
}

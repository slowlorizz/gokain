package wheel

type Wheel struct {
	/*
		A Type thats a "circular one way linked list"
		the next Item for the last ist the first.

		--> Its the Base-Type for "Gear" and "Seed"
	*/

	Root    *Item    // First Item
	Current Iterator // Current Item
}

type Iterator struct {
	/*
		A Wheel Iterator, that references a Wheel Item to iterate over a wheel
	*/
	Item *Item
	Root *Item
}

func New(from *[]string) *Wheel {

	whl := &Wheel{}

	var p *Item

	for _, v := range *from {
		itm := &Item{Char: v}

		if p == nil {
			whl.Root = itm
		} else {
			p.Next = itm
		}

		p = itm
	}

	p.Next = whl.Root
	whl.Current = Iterator{Item: whl.Root, Root: whl.Root}

	return whl
}

func (whl *Wheel) NewIterator() *Iterator {
	return &Iterator{Item: whl.Root, Root: whl.Root}
}

func (whl *Wheel) GetIteratorAt(start int) *Iterator {
	it := whl.NewIterator()

	for i := 0; i < start; i++ {
		it.Shift()
	}

	return it
}

func (it *Iterator) AtRoot() bool {
	return (it.Item.Char == it.Root.Char)
}

func (it *Iterator) AtEnd() bool {
	return (it.Item.Next.Char == it.Root.Char)
}

func (it *Iterator) Shift() bool {
	it.Item = it.Item.Next
	return it.AtRoot()
}

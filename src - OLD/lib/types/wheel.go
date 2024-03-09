package types

import "fmt"

type Item struct {
	Value rune
	Next  *Item
	Index int
}

func (itm *Item) Print(stopOn int) {
	fmt.Printf("[%d]=(%s) --> ", itm.Index, string(itm.Value))

	if itm.Index == stopOn {
		fmt.Println("\n\n\n")
		return
	}

	if itm.Next != nil || itm.Index != stopOn {
		itm.Next.Print(stopOn)
	}
}

type Wheel struct {
	First   *Item
	Last    *Item
	Current *Item
	Lenght  int
}

func New_Wheel(from string) Wheel {
	lttr := []rune(from)
	whl := Wheel{Lenght: len(from)}

	var prev *Item

	for i, r := range lttr {
		itm := &Item{Value: r, Index: i}

		if prev != nil {
			prev.Next = itm
		}

		prev = itm

		if i == 0 {
			whl.First = itm
		} else if i == (len(lttr) - 1) {
			whl.Last = itm
		}
	}

	whl.Last.Next = whl.First
	whl.Current = whl.First

	return whl
}

func (whl *Wheel) Value() rune {
	return whl.Current.Value
}

func (whl *Wheel) Turn() {
	whl.Current = whl.Current.Next
}

func (whl *Wheel) OnFirst() bool {
	return (whl.Current.Index == 0)
}

func (whl *Wheel) OnLast() bool {
	return (whl.Current.Index == (whl.Lenght - 1))
}

func (whl *Wheel) PrintItems(id *int) {
	fmt.Printf("\n\n===========[%d]=================\n", *id)
	whl.First.Print(91)
}

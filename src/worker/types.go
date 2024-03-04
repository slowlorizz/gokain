package gkwrk

import "strings"

type CircNode struct {
	Value string
	Prev  *CircNode
	Next  *CircNode
}

type CircList struct {
	First *CircNode
	Nodes []*CircNode
}

func NewCircList(from string) *CircList {
	letters := strings.Split(from, "")
	clist := CircList{}

	clist.Nodes = make([]*CircNode, 0)

	prev := &CircNode{}
	for _, s := range letters {
		nd := &CircNode{Value: s, Prev: prev}

		nd.Prev.Next = nd

		clist.Nodes = append(clist.Nodes, nd)
	}

	if len(clist.Nodes) > 1 {
		clist.Nodes[0].Prev = clist.Nodes[len(clist.Nodes)-1]
		clist.Nodes[len(clist.Nodes)-1].Next = clist.Nodes[0]
	}

	return &clist
}

type Wheel struct {
	Items   *CircList
	Current *CircNode
}

func NewWheel(from string) *Wheel {
	whl := Wheel{Items: NewCircList(from)}
	whl.Current = whl.Items.First

	return &whl
}

func (whl *Wheel) Reset() {
	whl.Current = whl.Items.First
}

func (whl *Wheel) Turn() {
	whl.Current = whl.Current.Next
}

func (whl *Wheel) Value() *string {
	return &whl.Current.Value
}

func (whl *Wheel) IsAtFirst() bool {
	return (whl.Current == whl.Items.First)
}

func (whl *Wheel) IsAtLast() bool {
	return (whl.Current == whl.Items.First.Prev)
}

type Combination struct {
	from_str string
	Wheels   []*Wheel
}

func NewCombination(from string, start_len int) *Combination {
	combo := Combination{from_str: from}
	combo.Wheels = make([]*Wheel, 0)

	for i := 0; i < start_len; i++ {
		combo.Wheels = append(combo.Wheels, NewWheel(combo.from_str))
	}

	return &combo
}

func (combo *Combination) String() *string {
	str := ""

	for _, w := range combo.Wheels {
		str += w.Current.Value
	}

	return &str
}

func (combo *Combination) Next() bool {
	turnNext := false
	for i, w := range combo.Wheels {
		turnNext = w.IsAtLast()

		w.Turn()

		if turnNext && i == (len(combo.Wheels)-1) {
			combo.Wheels = append(combo.Wheels, NewWheel(combo.from_str))

			break
		}

		if !turnNext {
			break
		}
	}

	return turnNext
}

type ResultPair struct {
	PlainText string
	Computed  string
}

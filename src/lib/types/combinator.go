package types

import (
	"fmt"
)

type Combinator struct {
	Wheel   Wheel
	Next    *Combinator
	fromStr string
}

func New_Combinator(from string) Combinator {
	return Combinator{Wheel: New_Wheel(from), fromStr: from}
}

func (this *Combinator) Turn(seeds *Wheel) {
	if this.Wheel.OnLast() {
		this.Wheel.Turn()

		if this.Next == nil {
			if seeds.OnLast() {
				nxt := New_Combinator(this.fromStr)
				this.Next = &nxt
			}

			seeds.Turn()

		} else {
			this.Next.Turn(seeds)
		}
	} else {
		this.Wheel.Turn()
	}
}

func (this *Combinator) GetCombo(str *string) {
	if this.Next != nil {
		this.Next.GetCombo(str)
	}

	*str += string(this.Wheel.Current.Value)

	return
}

func (this *Combinator) PrintWheel(id *int) {
	fmt.Println("---------------------------")
	*id += 1
	this.Wheel.PrintItems(id)

	if this.Next != nil {
		this.Next.PrintWheel(id)
	}
}

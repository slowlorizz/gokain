package combination

import (
	"fmt"

	"github.com/slowlorizz/gokain/worker/src/thread/combination/wheel"
)

type Gear struct {
	/*
		Type for a Wheel-Iterator, that has peers
	*/

	Itr  wheel.Iterator
	Next *Gear
	Seed *wheel.Iterator
}

func (G *Gear) Turn(t bool) string {
	if t {
		if G.Itr.Shift() {
			if G.Next == nil {
				if G.Seed.Shift() {
					G.Next = &Gear{Seed: G.Seed, Itr: wheel.Iterator{Item: G.Itr.Root, Root: G.Itr.Root}}
					return fmt.Sprintf("%s%s", G.Itr.Item.Char, G.Next.Itr.Item.Char)
				} else {
					return G.Itr.Item.Char
				}
			} else {
				return fmt.Sprintf("%s%s", G.Itr.Item.Char, G.Next.Turn(true))
			}
		}
	}

	if G.Next == nil {
		return G.Itr.Item.Char
	} else {
		return fmt.Sprintf("%s%s", G.Itr.Item.Char, G.Next.Turn(false))
	}
}

package combination

import (
	"github.com/slowlorizz/gokain/worker/src/thread/combination/wheel"
)

type Gear struct {
	/*
		Type for a Wheel-Iterator, that has peers
	*/

	Itr  *wheel.Iterator
	Next *Gear
}

func (G *Gear) Turn(seed *wheel.Iterator) string {
	if G.Itr.Shift() {
		if G.Next == nil {
			if seed.Shift() {
				G.Next = &Gear{Itr: &wheel.Iterator{Item: G.Itr.Root, Root: G.Itr.Root}}
				return G.Itr.Item.Char + G.Next.Itr.Item.Char
			}
		} else {
			return G.Itr.Item.Char + G.Next.Turn(seed)
		}
	}

	return G.Itr.Item.Char
}

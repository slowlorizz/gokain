package thread

import (
	"github.com/slowlorizz/gokain/worker/src/thread/combination"
	"github.com/slowlorizz/gokain/worker/src/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/src/ui"
)

type (
	ResultPair struct {
		PlainText string
		Hash      string
	}

	Signal = struct{}

	Thread struct {
		ID          int
		Combination combination.Combination
		CipherText  string
		UiC         *ui.ThreadComponent
		FoundCH     chan ResultPair
		JoinCH      chan Signal
	}
)

var FOUND_CH (chan ResultPair) = make(chan ResultPair)
var JOIN_CH (chan Signal) = make(chan Signal)

var JOINED bool = false

func New(id int, ct string, seeds []string, ht combination.HashType, cs charset.CharSet, foundCH chan ResultPair, joinCH chan Signal) *Thread {
	return &Thread{
		ID:          id,
		Combination: *combination.New(seeds, cs, ht),
		FoundCH:     foundCH,
		JoinCH:      joinCH,
		UiC:         ui.New_ThreadComponent(id, ht),
		CipherText:  ct,
	}
}

func (T *Thread) Start() {
	for {
		select {
		case <-T.JoinCH:
			return
		default:
			pt, hsh := T.Combination.Next()

			T.UiC.HashC.Text = hsh
			T.UiC.PtC.Text = pt

			if hsh == T.CipherText {
				T.FoundCH <- ResultPair{Hash: hsh, PlainText: pt}
				T.UiC.SetStyleFound()
				return
			}
		}
	}
}

func StopAll() {
	if !JOINED {
		close(JOIN_CH)
		JOINED = true
	}
}

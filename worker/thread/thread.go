package thread

import (
	"github.com/slowlorizz/gokain/worker/thread/combination"
	"github.com/slowlorizz/gokain/worker/thread/combination/charset"
	"github.com/slowlorizz/gokain/worker/ui"
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
		UiComponent *ui.ThreadComponent
		FoundCH     chan<- ResultPair
		JoinCH      <-chan Signal
	}
)

var FOUND_CH (chan ResultPair) = make(chan ResultPair)
var JOIN_CH (chan Signal) = make(chan Signal)

func New(id int, ct string, seeds []string, ht combination.HashType, cs charset.CharSet, foundCH chan<- ResultPair, joinCH <-chan Signal) *Thread {
	// charset.CharSet{Numbers: true, LowerCase: true, UpperCase: true, Special: true, Ext_LC: false, Ext_UC: false, Ext_Spc: false}
	return &Thread{
		Combination: *combination.New(seeds, cs, ht),
		FoundCH:     foundCH,
		JoinCH:      joinCH,
		UiComponent: ui.New_ThreadComponent(id),
		CipherText:  ct,
	}
}

func (T *Thread) Start() {
	for {
		select {
		case <-T.JoinCH:
			return
		default:
			pt, hsh := T.Combination.Cycle()

			T.UiComponent.HashC.SetText(hsh)
			T.UiComponent.PtC.SetText(pt)

			if hsh == T.CipherText {
				T.FoundCH <- ResultPair{Hash: hsh, PlainText: pt}
				T.UiComponent.SetStyleFound()
				return
			}
		}
	}
}

func StopAll() {
	JOIN_CH <- Signal{}
	close(JOIN_CH)
}

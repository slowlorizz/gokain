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
		Combination combination.Combination
		CipherText  string
		UiComponent *ui.ThreadComponent
		FoundCH     chan<- ResultPair
		JoinCH      <-chan Signal
	}
)

func New(ct string, seeds []string, ht combination.HashType, cs charset.CharSet, foundCH chan<- ResultPair, joinCH <-chan Signal, uiComp *ui.ThreadComponent) *Thread {
	// charset.CharSet{Numbers: true, LowerCase: true, UpperCase: true, Special: true, Ext_LC: false, Ext_UC: false, Ext_Spc: false}
	th := &Thread{
		Combination: *combination.New(seeds, cs, ht),
		FoundCH:     foundCH,
		JoinCH:      joinCH,
		UiComponent: uiComp,
		CipherText:  ct,
	}

	return th
}

func (T *Thread) Start() {
	for {
		select {
		case <-T.JoinCH:
			return
		default:
			pt, hsh := T.Combination.Cycle()

			if hsh == T.CipherText {
				T.FoundCH <- ResultPair{Hash: hsh, PlainText: pt}
			}
		}
	}
}

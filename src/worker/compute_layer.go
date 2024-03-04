package gkwrk

type ComputeLayer struct {
	charMap       string
	seeds         string
	maxRoutines   int // How many concurent routines are allowed to run
	outCh         *(chan *ResultPair)
	stopCh        *(chan bool)
	Action        func(v *string) *ResultPair // --> the Action which th Combinators are going to run
	aciveRoutines int
	combinators   []*Combinator
}

func (cmp *ComputeLayer) Start() {
	if len(cmp.seeds) < cmp.maxRoutines {
		cmp.aciveRoutines = len(cmp.seeds)
	} else {
		cmp.aciveRoutines = cmp.maxRoutines
	}

	cSeeds := make([]string, cmp.aciveRoutines)
	cmp.combinators = make([]*Combinator, 0)

	for i := 0; i < cmp.aciveRoutines; i++ {
		cSeeds[i] = ""
	}

	for i := 0; i < len(cmp.seeds); i++ {
		cSeeds[(i % cmp.aciveRoutines)] = string(cmp.seeds[i])
	}

	for _, v := range cSeeds {
		cmb := cmp.NewCombinator(cmp.charMap, v, 1)
		go cmb.Run()
	}
}

type Combinator struct {
	charMap     *Wheel
	combination *Combination
	comboStr    string
	seed        *Wheel
	outCh       *(chan *ResultPair)
	stopCh      *(chan bool)
	Action      func(v *string) *ResultPair
}

func (cmp *ComputeLayer) NewCombinator(chr_map string, seeds string, start_len int) *Combinator {
	cmb := &Combinator{
		charMap:     NewWheel(chr_map),
		combination: NewCombination(chr_map, start_len),
		comboStr:    "",
		seed:        NewWheel(seeds),
		outCh:       cmp.outCh,
		stopCh:      cmp.stopCh,
		Action:      cmp.Action,
	}

	cmp.combinators = append(cmp.combinators, cmb)

	return cmb
}

func (cmb *Combinator) Combine() {
	finish_len := cmb.combination.Next() // finished_len is true when all combinations of the length where run and a new wheel was added --> then shift to new seed
	cmb.comboStr = *cmb.seed.Value()
	cmb.comboStr += *cmb.combination.String()
	*cmb.outCh <- cmb.Action(&cmb.comboStr)

	if finish_len {
		cmb.seed.Turn()
	}
}

func (cmb *Combinator) Run() {
	cmb.Combine()

	for {
		select {
		case <-*cmb.stopCh:
			return
		default:
			cmb.Combine()
		}
	}
}

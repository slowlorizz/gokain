package types

type Combination struct {
	First   Combinator
	CharMap string
	Seed    Wheel
}

func GenCharMap(nums bool, lower bool, upper bool, special bool) string {
	chrMap := ""

	if nums {
		chrMap += "0123456789"
	}

	if lower {
		chrMap += "abcdefghijklmnopqrstuvwxyz"
	}

	if upper {
		chrMap += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if special {
		chrMap += "<>;,:.-_${}[]!^~`'?=()/&%*\"#@|"
	}

	return chrMap
}

func New_Combination(seed_str string, nums bool, lower bool, upper bool, special bool) Combination {
	combo := Combination{CharMap: GenCharMap(nums, lower, upper, special), Seed: New_Wheel(seed_str)}
	combo.First = New_Combinator(combo.CharMap)

	return combo
}

func (this *Combination) Next() {
	this.First.Turn(&this.Seed)
}

func (this *Combination) Cycle() string {
	val := string(this.Seed.Current.Value)
	this.First.GetCombo(&val)

	this.Next()

	return val
}

func (this *Combination) PrintCombinators() {
	id := 0
	this.First.PrintWheel(&id)
	//this.First.Wheel.PrintItems(&id)
}

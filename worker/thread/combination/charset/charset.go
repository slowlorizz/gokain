package charset

import "strings"

type CharSet struct {
	Numbers   bool
	LowerCase bool
	UpperCase bool
	Special   bool
	Ext_LC    bool // extended lowercase --> includes some region specific characters
	Ext_UC    bool // extended uppercase --> includes some region specific characters
	Ext_Spc   bool // extended special-characters --> includes some region specific characters
	Chars     []string
}

const (
	NUMBERS    string = "0123456789"                       // 10
	LOWER_CASE string = "abcdefghijklmnopqrstuvwxyz"       // 26
	UPPER_CASE string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"       // 26
	SPECIAL    string = "+!$-_*%&/=?~:.,;|@#\"'(){}[]<>\\" // 30
	EXT_LC     string = "äöüéàè"                           // 6
	EXT_UC     string = "ÄÖÜÉÀÈ"                           // 6
	EXT_SPC    string = "°§£^`¨¬¢€"                        // 9
)

func (cs *CharSet) Add(str string) {
	cs.Chars = append(cs.Chars, strings.Split(str, "")...)
}

func (cs *CharSet) Build() {
	cs.Chars = make([]string, 0, 113)

	if cs.Numbers {
		cs.Add(NUMBERS)
	}

	if cs.LowerCase {
		cs.Add(LOWER_CASE)
	}

	if cs.Ext_LC {
		cs.Add(EXT_LC)
	}

	if cs.UpperCase {
		cs.Add(UPPER_CASE)
	}

	if cs.Ext_UC {
		cs.Add(EXT_UC)
	}

	if cs.Special {
		cs.Add(SPECIAL)
	}

	if cs.Ext_Spc {
		cs.Add(EXT_SPC)
	}
}

func New(nums bool, lc bool, uc bool, spc bool, ext_lc bool, ext_uc bool, ext_spc bool) *CharSet {
	cs := CharSet{
		Numbers:   nums,
		LowerCase: lc,
		UpperCase: uc,
		Special:   spc,
		Ext_LC:    ext_lc,
		Ext_UC:    ext_uc,
		Ext_Spc:   ext_spc,
	}

	cs.Build()

	return &cs
}

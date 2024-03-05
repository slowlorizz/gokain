package feat

import (
	tui "gokain/tui"
	types "gokain/types"
	utils "gokain/utils"
)

func CrackHash(stopCH chan bool, foundCH chan types.ResultPair /*Return Found Plaintext*/, hash string, hashType types.HashType, seed string, ui_component *tui.ThreadComponent) {
	combo := types.New_Combination(seed, true, true, true, true)

check_loop:
	for {
		select {
		case <-stopCH:
			break check_loop

		default:
			combo_val := combo.Cycle()
			cHash := utils.ComputeHash(combo_val, hashType)

			ui_component.SetHashText(cHash)
			ui_component.SetPlainText(combo_val)

			if cHash == hash {
				ui_component.SetBordersGreen()
				foundCH <- types.ResultPair{PlainText: combo_val, Hash: cHash}
				break check_loop
			}
		}
	}
}

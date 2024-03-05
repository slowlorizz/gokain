package feat

import (
	types "gokain/types"
	utils "gokain/utils"

	logs "github.com/slowlorizz/gokain-logs"
)

func CrackHash(stopCH chan bool, foundCH chan types.ResultPair /*Return Found Plaintext*/, hash string, hashType types.HashType, seed string) {
	logs.Debug("Starting hash-Crack...")
	combo := types.New_Combination(seed, true, true, true, true)

check_loop:
	for {
		select {
		case v := <-stopCH:
			stopCH <- v

		default:
			combo_val := combo.Cycle()
			cHash := utils.ComputeHash(combo_val, hashType)

			if cHash == hash {
				foundCH <- types.ResultPair{PlainText: combo_val, Hash: cHash}
				break check_loop
			} else {
				logs.Debug(combo_val, "Hash", cHash)
			}
		}
	}
}

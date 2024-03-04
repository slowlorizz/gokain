package main

import (
	gkwrk "gokain/worker"

	logs "github.com/slowlorizz/gokain-logs"
)

//  go mod edit -replace gokain/logs=../lib/logs

func main() {
	args := Args{}
	args.Load()

	logs.Handler.Start(args.verbose) // starts the async Log-Handler

	wrk := gkwrk.Worker{
		CipherText: args.hash,
		Nums:       true,
		Lower:      true,
		Upper:      true,
		Special:    true,
	}

	wrk.Start(5)
	wrk.Join()
	logs.Handler.Join() // Prints the last received Log-Message before program-end
}

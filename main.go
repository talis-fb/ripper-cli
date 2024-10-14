package main

import (
	"fmt"
	"os"

	"github.com/talis-fb/ripper-cli/internal"
)

func main() {
	cliEntry, err := internal.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	startPrefix := []byte(cliEntry.StartPrefix)
	endPrefix := []byte(cliEntry.EndPrefix)

	errRipper := internal.Ripper(startPrefix, endPrefix)
	if errRipper != nil {
		fmt.Println(errRipper)
		os.Exit(1)
		return
	}
}

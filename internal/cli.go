package internal

import (
	"flag"
	"fmt"
)

type CliEntry struct {
	StartPrefix string
	EndPrefix   string
}

func Init() (CliEntry, error) {
	// version := flag.Bool("v", false, "Enable verbose mode")
	flag.Parse()

	params := flag.Args()

	if len(params) < 2 {
		PrintHelp()
		return CliEntry{}, fmt.Errorf("not enough arguments")
	}

	firstParam := params[0]
	secondParam := params[1]

	return CliEntry{firstParam, secondParam}, nil
}

func PrintHelp() {
	fmt.Println("Usage: ripper <pattern1> <pattern2>")
	fmt.Println("This tool expects exactly two arguments which are valid regex patterns.")
}

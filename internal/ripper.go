package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Ripper(startPrefix, endPrefix []byte) error {
	reader := bufio.NewReader(os.Stdin)

	inSlice := false
	for {
		buffer, err := reader.ReadBytes('\n')
		if err != nil {
			return err
		}

		if bytes.Contains(buffer, startPrefix) {
			inSlice = true
		}

		if bytes.Contains(buffer, endPrefix) {
			inSlice = false
		}

		if inSlice {
			fmt.Print(string(buffer))
		}
	}

	return nil
}

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	data := ""

	if len(os.Args) > 1 {
		// get the supplied args
		for _, s := range os.Args[1:] {
			if len(s) < 2 {
				s = fmt.Sprintf("%02v", s)
			}
			data += s
		}

	}

	if len(data) == 0 {
		// no data so trying stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			data = s.Text()
		}
	}

	data = strings.Replace(data, " ", "", -1)
	data = strings.Replace(data, ",", "", -1)
	data = strings.Replace(data, "0x", "", -1)

	// decode the hex to bytes
	b, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}

	// write the bytes out as a string
	fmt.Printf("%s\n", b)
}

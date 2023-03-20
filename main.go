package main

import (
	"fmt"
	"os"
)

/*
	Goal: Spam masker
	Input: Welcome to https://twitter.com/TheNishuGoel profile.
	Output: Welcome to https://************************ profile.
*/

func main() {

	const (
		linkType       = "https://" // link pattern to mask
		linkTypeLength = len(linkType)
		mask           = '*'
	)

	// Prompt
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Enter something to mask :)")
		return
	}

	var (
		input = args[0]               // get input to mask
		size  = len(input)            // get input length to prepare a byte buffer
		buf   = make([]byte, 0, size) // byte buffer to store input

		in bool // bool to check if a link is found
	)

	for i := 0; i < size; i++ {

		// find a link pattern to mask
		if len(input[i:]) > linkTypeLength && linkType == input[i:i+linkTypeLength] {
			in = true

			// keep 'https' the same in the masked link
			buf = append(buf, linkType...)
			i += linkTypeLength
		}

		c := input[i]

		// avoid masking keywords after link
		switch c {
		case ' ', '\t', '\n':
			in = false
		}
		if in {
			c = mask
		}

		buf = append(buf, c)
	}
	fmt.Println(string(buf))
}

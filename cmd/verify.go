package main

import (
	"fmt"

	"github.com/jjshanks/cryptopals/internal/pkg/base64"
	"github.com/jjshanks/cryptopals/internal/pkg/hex"
)

func main() {
	// Set 1 Challenge 1
	byteInput, err := hex.Decode("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		fmt.Printf("Failed to convert input to bytes for set 1 challenge 1: %v\n", err)
		return
	}
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Printf("Set 1 Challenge 1 match? %v\n", expectedOutput == base64.Encode(byteInput))
}

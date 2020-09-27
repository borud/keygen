package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

// Options contains the command line options
type Options struct {
	KeySize int `short:"s" long:"size" default:"256" description:"Key size in bits.  Must be a multiple of 8."`
}

var opt Options
var parser = flags.NewParser(&opt, flags.Default)

func main() {
	_, err := parser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
		log.Fatalf("Internal error: %v", err)
	}

	// Check that this is a reasonable keysize
	if (opt.KeySize % 8) != 0 {
		log.Fatalf("Number of bits has to be a multiple of 8")
	}

	bufLen := opt.KeySize / 8
	keyBuffer := make([]byte, bufLen)

	n, err := rand.Read(keyBuffer)
	if err != nil {
		log.Fatalf("Unable to generate key: %v", err)
	}

	// This shouldn't happen but it doesn't hurt to be careful
	if n < bufLen {
		log.Fatalf("Wrote too few bytes: %v", err)
	}

	fmt.Printf("\nKey length is %d bits (%d bytes)\n\n", opt.KeySize, bufLen)

	// Print Base64
	keyStr := base64.StdEncoding.EncodeToString(keyBuffer)
	fmt.Printf(" Base64  : \"%s\"\n", keyStr)

	// Print Hex string
	fmt.Printf(" Hex     : \"%x\"\n", keyStr)

	// Print C-style byte array
	fmt.Printf(" C-array : {")
	var pieces = make([]string, bufLen)
	for n, b := range keyBuffer {
		pieces[n] = fmt.Sprintf("0x%x", b)
	}
	fmt.Printf("%s}\n", strings.Join(pieces, ","))
	fmt.Println()
}

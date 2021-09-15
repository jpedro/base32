package main

import (
	"bufio"
	"encoding/base32"
	"fmt"
	"os"
)

var (
	VERSION = "v0.1.0"
	USAGE   = `SYNOPSIS
    Base32 encodes and decodes payloads

USAGE
    base32 encode [PAYLOAD]     # Encodes PAYLOAD or what's in the STDIN
    base32 decode [PAYLOAD]     # Decodes PAYLOAD or what's in the STDIN
    base32 help                 # Shows this help
    base32 version              # Shows the current version
`
)

func main() {
	command := "help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "help", "--help":
		fmt.Print(USAGE)
		return

	case "version", "--version":
		fmt.Println(VERSION)
		return

	case "encode", "enc", "e", "--encode", "-e":
		payload := getPayload(command, os.Args)
		encoded := encode(payload)
		fmt.Println(encoded)

	case "decode", "dec", "d", "--decode", "-d":
		payload := getPayload(command, os.Args)
		decoded := decode(payload)
		fmt.Println(decoded)

	default:
		fmt.Fprintf(os.Stderr, "Error: Command '%s' not found.\n", command)
		fmt.Fprintf(os.Stderr, "Run 'base32 help' to check available commands.\n")
		os.Exit(1)
	}
}

func getPayload(command string, args []string) string {
	payload := ""
	if len(args) > 2 {
		return args[2]
	}
	payload = getStdin(command)
	return payload
}

// Encodes payload into a base32 string
func encode(payload string) string {
	bytes := []byte(payload)
	encoded := base32.StdEncoding.EncodeToString(bytes)

	return encoded
}

// Decodes payload from a base32 encoded string
func decode(encoded string) string {
	decoded, err := base32.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	return string(decoded)
}

func getStdin(command string) string {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe == 0 {
		fmt.Printf("Enter the text to %s and finish with Ctrl+D:\n", command)
	}

	payload := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		payload = scanner.Text()
	}

	return payload
}

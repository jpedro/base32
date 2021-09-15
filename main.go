package main

import (
	"bufio"
	"encoding/base32"
	"fmt"
	"os"
)

func main() {
	command := "help"
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "help", "--help":
		fmt.Printf("USAGE %s\n", "base32")
		return

	case "version", "--version":
		fmt.Printf("v%s\n", "0.1.0")
		return

	case "encode", "enc", "e", "--encode", "-e":
		payload := getPayload(command, os.Args)
		encoded := encode(payload)
		fmt.Println(encoded)

	case "decode", "dec", "d", "--decode", "-d":
		payload := getPayload(command, os.Args)
		decoded := decode(payload)
		fmt.Println(decoded)
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

func encode(payload string) string {
	bytes := []byte(payload)
	encoded := base32.StdEncoding.EncodeToString(bytes)

	return encoded
}

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

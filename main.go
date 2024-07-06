package main

import (
	"buzzdl/utils"
	"flag"
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	var (
		format string
		url string
		output string
	)

	flag.StringVar(&format, "format", "mp4", "output format")
	flag.StringVar(&format, "f", "mp4", "output format(shorthand)")

	flag.StringVar(&url, "url", "", "bsocial url")
	flag.StringVar(&url, "u", "", "bsocial url(shorthand)")

	flag.StringVar(&output, "output", "./", "output location")
	flag.StringVar(&output, "o", "./", "output location(shorthand)")

	flag.Parse()

	if url == "" {
		rl, err := readline.New("Bsocial URL?: ")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		url, err = rl.Readline()
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

	if url == "" {
		fmt.Println("Error: no url provided")
		return
	}

	if format == "" {
		fmt.Println("Error: A format is required")
		return
	}

	// Convert to lowercase
	format = strings.ToLower(format)
	url = strings.ToLower(url)

	utils.StartProgram(url, format, output)
}

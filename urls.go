package urls

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func Run(process func(string)) {
	if len(os.Args) == 2 {
		process(os.Args[1])
	} else {
		// if argument isn't present, read from stdin, to allow piped usage
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			process(scanner.Text())
		}
	}
}

func Unescape(text string) {
	escaped, err := url.QueryUnescape(text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(escaped)
}

func Escape(text string) {
	escaped := url.QueryEscape(text)
	fmt.Println(escaped)
}

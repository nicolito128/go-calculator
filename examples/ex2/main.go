package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nicolito128/go-calculator"
)

func main() {
	fmt.Println("Calculator!")
	fmt.Println("-- Enter a new inline-operation and press enter.")
	fmt.Println("-- If you want to end the calculator enter 'e' or 'end' and press enter.")
	scanner := bufio.NewScanner(os.Stdin)

	var line string
	for {
		scanner.Scan()
		line = scanner.Text()

		if strings.ToLower(line) == "e" || strings.ToLower(line) == "end" {
			fmt.Println("See you later! -- PROGRAM END")
			break
		}

		result, err := calculator.Resolve(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}

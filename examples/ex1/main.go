package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nicolito128/go-calculator"
)

func main() {
	fmt.Println("Calculator!")
	fmt.Println("-- Enter an operation")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	result, err := calculator.Resolve(scanner.Text())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

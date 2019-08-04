package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func scanandreadandjoin() {
	input := "foo bar baz" // or os.Stdin
	var buf bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
		buf.Write(scanner.Bytes())
	}

	output := buf.Bytes()
	fmt.Println(output, output[0])
}

func main() {
	scanandreadandjoin()
}

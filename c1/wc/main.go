package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) (int, error) {
	scanner := bufio.NewScanner(r)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countLines && countBytes {
		return -1, errors.New("only one option expected")
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanLines) // ScanLines is already default, but it's better to be explicit about things sometimes
	}
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc, nil
}

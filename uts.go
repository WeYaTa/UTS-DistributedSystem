package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

		words := strings.Split(scanner.Text(), " ")
		numberofword := make(map[string]int)

		count(words, numberofword)

	}
}

func count(a []string, now map[string]int) {
	for _, i := range a {
		if _, ok := now[i]; ok {
			now[i]++
		} else {
			now[i] = 1
		}
	}
	for key, value := range now {
		fmt.Printf("%s : %d\n", key, value)
	}
}

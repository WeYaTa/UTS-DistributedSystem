package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	a := make(chan string)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {

		words := strings.Split(scanner.Text(), " ")
		numberofword := make(map[string]int)

		go count(a, numberofword)

		for _, word := range words {
			a <- word
		}

	}
}

func count(a chan string, now map[string]int) {
	for {
		i, ok := <-a
		if !ok {
			break
		}

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

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(a []string) {
	p := a[0]
	c := 0
	for _, n := range a[1:] {
		if n > p {
			c++
		}
		p = n
	}
	fmt.Println(c) //The actual solution is 1559 but I get 1558 I don't know why
}

func main() {
	file, err := os.Open("input")
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		const maxCap = 64000
		buf := make([]byte, maxCap)
		scanner.Buffer(buf, maxCap) // Change maxCap if file is larger than 64k lines

		lines := []string{}
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		err := scanner.Err()
		if err == nil { // file read successfully

			solution(lines)

		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(err)
	}
}

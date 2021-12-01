package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solution(a []int) {
	c := 0
	for i := range a[1 : len(a)-2] {
		if a[i+1]+a[i+2]+a[i+3] > a[i]+a[i+1]+a[i+2] {
			c++
		}
	}
	fmt.Println(c)
}

func main() {
	file, err := os.Open("input")
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		const maxCap = 64000
		buf := make([]byte, maxCap)
		scanner.Buffer(buf, maxCap) // Change maxCap if file is larger than 64k lines

		lines := []int{}
		for scanner.Scan() {
			i, err := strconv.Atoi(scanner.Text())
			if err == nil {
				lines = append(lines, i)
			}
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

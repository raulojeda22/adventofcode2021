package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(a []string) {
	h := 0
	d := 0
	for _, v := range a {
		o := strings.Split(v, " ")
		n, err := strconv.Atoi(o[1])
		if err == nil {
			if o[0] == "forward" {
				h = h + n
			} else if o[0] == "down" {
				d = d + n
			} else if o[0] == "up" {
				d = d - n
			}
		}
	}
	fmt.Println(h, d)
	fmt.Println(h * d)
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

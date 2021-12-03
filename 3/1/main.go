package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(a []string) { // 10:31
	d := strings.Split(a[0], "")
	g := ""
	e := ""
	for i := 0; i < len(d); i++ {
		o := 0
		z := 0
		for _, v := range a {
			if v[i:i+1] == "1" {
				o++
			} else {
				z++
			}
		}
		if o >= z {
			g = g + "1"
			e = e + "0"
		} else {
			g = g + "0"
			e = e + "1"
		}
	}
	x, _ := strconv.ParseInt(g, 2, 64)
	y, _ := strconv.ParseInt(e, 2, 64)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x * y)
}

func main() {
	file, err := os.Open("input")
	_, _ = strconv.Atoi("9")
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solution(a []string) { // 10:31 -> 44:10
	d := strings.Split(a[0], "")
	g := ""
	e := ""
	p := a
	r := []string{}
	var s int64 = 1
	for j := 0; j < 2; j++ {
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
				t := []string{}
				for _, v := range a {
					if j == 0 {
						if v[i:i+1] == "1" {
							t = append(t, v)
						}
					} else {
						if v[i:i+1] == "0" {
							t = append(t, v)
						}
					}
				}
				a = t
				g = g + "1"
				e = e + "0"
			} else {
				t := []string{}
				for _, v := range a {
					if j == 0 {
						if v[i:i+1] == "0" {
							t = append(t, v)
						}
					} else {
						if v[i:i+1] == "1" {
							t = append(t, v)
						}
					}
				}
				a = t
				g = g + "0"
				e = e + "1"
			}
			if len(a) == 1 {
				break
			}
		}
		fmt.Println(a[0])
		r = append(r, a[0])
		a = p
	}
	for _, v := range r {
		re, _ := strconv.ParseInt(v, 2, 64)
		fmt.Println(re)
		s = s * re
	}
	fmt.Println(s)
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

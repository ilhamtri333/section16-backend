package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var data = make(map[rune]int)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()
		frequency(inputString)
		time.Sleep(500 * time.Millisecond)
		showDataMap()
	}
}

func frequency(inputString string) {
	for _, k := range inputString {
		go frequencyChar(k, inputString)
	}
}

func frequencyChar(c rune, inputString string) {
	var Exist = false
	for k, _ := range data {
		if k == c {
			Exist = true
			break
		}
	}
	if !Exist {
		countString(c, inputString)
	}
}

func countString(c rune, inputString string) {
	var count int
	for _, i := range inputString {
		if i == c {
			count = count + 1
		}
	}
	data[c] = count
}

func showDataMap() {
	for i, val := range data {
		// var str = strconv.Itoa(val)
		character := string(i)
		fmt.Println(character, val)
	}
}

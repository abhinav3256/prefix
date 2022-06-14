package main

import (
	"fmt"
	"strings"
)

func main() {

	var input = []string{"pay", "attention", "practice", "attend"}
	var check int

	//var pre = "pre"
	prefArr := strings.Split("at", "")
	//fmt.Println(prefArr)
	var count int
	for i := 0; i < len(input); i++ {
		inputArr := strings.Split(input[i], "")
		//fmt.Println(inputArr)
		check = 0
		for j := 0; j < len(prefArr); j++ {
			if inputArr[j] == prefArr[j] {

				check = check + 1
			}
		}

		if check == len(prefArr) {
			count = count + 1
			fmt.Println(input[i])
		}

	}
	fmt.Println(count)

}

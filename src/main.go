package main

import (
	"fmt"
)

func main() {
	counts := make(map[string]int)
	for {
		var inputStr string
		fmt.Scanf("%s", &inputStr)
		if "exit" == inputStr {
			break
		} else {
			counts[inputStr]++
		}
	}
	for key, value := range counts {
		fmt.Printf("%s出现%d次 \n", key, value)
	}

}

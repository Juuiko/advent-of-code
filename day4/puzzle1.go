package main

import (
	"fmt"
	"strconv"
)

func main() {
	lowerBound, upperBound := 108457, 562041
	totalNums := 0
	for i:=lowerBound; i<=upperBound; i++ {
		n := strconv.Itoa(i)
		if n[0]==n[1]||n[1]==n[2]||n[2]==n[3]||n[3]==n[4]||n[4]==n[5] {
			if n[0]<=n[1]&&n[1]<=n[2]&&n[2]<=n[3]&&n[3]<=n[4]&&n[4]<=n[5]{
				totalNums++
			}
		}
	}
	fmt.Println(totalNums)
}

package main

import (
	"fmt"

	"techytechster.com/aoc-2021/d01/models"
)

func main() {
	sub := models.Submarine{}
	fmt.Println("---- Part One ----")
	count, err := sub.ComputeIncreasing("./static/part01.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
	fmt.Println("---- Part Two ----")
	count, err = sub.ComputeSlidingWindow("./static/part01.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(count)
	}
}

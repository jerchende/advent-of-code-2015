package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	input, _ := ioutil.ReadFile(pwd + "/day02/input.txt")
	lines := strings.Split(string(input), "\n")

	println(calculatePackageUsage(lines))
}

func calculatePackageUsage(input []string) (paper int, ribbon int) {
	for _, line := range input {
		p, r := forPackage(line)
		paper += p
		ribbon += r
	}
	return paper, ribbon
}

func forPackage(input string) (paper int, ribbon int) {
	var l, w, h int
	_, _ = fmt.Sscanf(input, "%dx%dx%d", &l, &w, &h)

	paper = 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	ribbon = min(2*l+2*w, 2*l+2*h, 2*h+2*w) + l*w*h
	return
}

func min(nums ...int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}

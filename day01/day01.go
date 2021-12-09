package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	input, _ := ioutil.ReadFile(pwd + "/day01/input.txt")

	println("floor:", calcFloor(input))
	println("basement:", calcBasement(input))
}

func calcBasement(input []byte) int {
	floor := 0
	for index, b := range input {
		switch b {
		case '(':
			floor++
		case ')':
			floor--
		}
		if floor == -1 {
			return index + 1
		}
	}
	return -1
}

func calcFloor(input []byte) int {
	return bytes.Count(input, []byte("(")) - bytes.Count(input, []byte(")"))
}

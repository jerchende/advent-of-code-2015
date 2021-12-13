package main

import (
	"io/ioutil"
	"os"
)

func main() {
	pwd, _ := os.Getwd()
	input, _ := ioutil.ReadFile(pwd + "/day03/input.txt")

	println(calculateHomesVisited(input, nil))
	i1, i2 := split(input)
	println(calculateHomesVisited(i1, i2))
}

func split(input []byte) (input1 []byte, input2 []byte) {
	for index, command := range input {
		if index%2 == 0 {
			input1 = append(input1, command)
		} else {
			input2 = append(input2, command)
		}
	}
	return
}

func calculateHomesVisited(input []byte, input2 []byte) (homes int) {
	var x, y uint16
	homesVisited := make(map[uint32]uint8)
	homesVisited[0]++
	for _, command := range input {
		moveCursor(command, &x, &y)
		homesVisited[uint32(x)<<16|uint32(y)]++
	}
	if input2 != nil {
		x, y = 0, 0
		for _, command := range input2 {
			moveCursor(command, &x, &y)
			homesVisited[uint32(x)<<16|uint32(y)]++
		}
	}
	return len(homesVisited)
}

func moveCursor(command byte, x *uint16, y *uint16) {
	switch command {
	case 'v':
		*y++
	case '^':
		*y--
	case '<':
		*x--
	case '>':
		*x++
	}
}

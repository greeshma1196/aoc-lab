package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer f.Close()

	r := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r = append(r, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return r
}

func aocDay1(input []string) (int, int) {

	var currPos int = 50
	var part1 int = 0
	var part2 int = 0
	for _, cmd := range input {
		dir := cmd[0]
		steps, _ := strconv.Atoi(cmd[1:])

		switch dir {
		case 'R':
			currPos, part2 = shfitR(currPos, int(steps), part2)
		case 'L':
			currPos, part2 = shiftL(currPos, int(steps), part2)
		default:
			fmt.Println("Unknown direction:", string(dir))
		}

		if currPos == 0 {
			part1++
		}

		fmt.Printf("Cmd: %s, Current Position: %d, Sub Counter: %d\n", cmd, currPos, part2)

	}

	part2 += part1
	return part1, part2
}

func shfitR(currPos int, steps int, subCounter int) (int, int) {
	newPos := currPos + steps

	if newPos > 100 {
		newPos, subCounter = continousShiftR(newPos, subCounter)
	}

	if newPos == 100 {
		newPos = 0
	}

	return newPos, subCounter
}

func continousShiftR(currPos int, subCounter int) (int, int) {
	if currPos > 100 {
		currPos = currPos - 100
		subCounter++
		return continousShiftR(currPos, subCounter)
	}
	return currPos, subCounter
}

func shiftL(currPos int, steps int, subCounter int) (int, int) {
	newPos := currPos - steps
	if newPos < 0 {
		if currPos != 0 {
			subCounter++
		}
		newPos, subCounter = continousShiftL(newPos, subCounter)
	}
	return newPos, subCounter
}

func continousShiftL(currPos int, subCounter int) (int, int) {
	if currPos < 0 {
		currPos = currPos + 100
		if currPos < 0 {
			subCounter++
		}
		return continousShiftL(currPos, subCounter)
	}
	return currPos, subCounter
}

func main() {
	r := readFile("inputs/input-1.txt")
	p1, p2 := aocDay1(r)
	fmt.Println("Result part 1:", p1)
	fmt.Println("Result part 2:", p2)
}

package solution_1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Calibrate2() int {
	numbers := map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	flag := false
	first := 0
	last := 0
	total := 0

	file, err := os.Open("solutions/day_1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		flag = false

		for key, val := range numbers {
			line = strings.ReplaceAll(line, key, strconv.Itoa(val))
		}

		for _, c := range line {
			if unicode.IsDigit(c) {
				if !flag {
					first = int(c - '0')
					flag = true
				}

				last = int(c - '0')
			}
		}

		

		total += first*10 + last
	}

	return total
}

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
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

		for i, val := range numbers {
			line = strings.ReplaceAll(line, val, strconv.Itoa(i+1))
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

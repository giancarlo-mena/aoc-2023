package solution_1

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func Calibrate() int {
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
		flag = false;

		for _, c := range line {
			if unicode.IsDigit(c) {
				if !flag {
					first = int(c - '0')
					flag = true
				} 

				last = int(c - '0')
			}
		}

		total += first * 10 + last
	}

	return total
}
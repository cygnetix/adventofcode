package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const seperator = "   "

func distance(list io.Reader) (int, error) {
	var list1, list2 []int
	var distance int

	scanner := bufio.NewScanner(list)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), seperator)

		left, err := strconv.Atoi(text[0])
		if err != nil {
			return 0, err
		}

		right, err := strconv.Atoi(text[1])
		if err != nil {
			return 0, err
		}

		list1 = append(list1, left)
		list2 = append(list2, right)
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	sort.Ints(list1)
	sort.Ints(list2)

	for i, val := range list1 {
		distance += int(math.Abs(float64(val - list2[i])))
	}

	return distance, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	d, err := distance(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("Total distance: %d", d)
}

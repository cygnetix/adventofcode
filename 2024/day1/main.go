package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const seperator = "   "

func parse(list io.Reader) (list1, list2 []int, err error) {
	scanner := bufio.NewScanner(list)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), seperator)

		left, err := strconv.Atoi(text[0])
		if err != nil {
			return nil, nil, err
		}

		right, err := strconv.Atoi(text[1])
		if err != nil {
			return nil, nil, err
		}

		list1 = append(list1, left)
		list2 = append(list2, right)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2, nil
}

func distance(list1, list2 []int) (int, error) {
	var distance int
	for i, val := range list1 {
		distance += int(math.Abs(float64(val - list2[i])))
	}

	return distance, nil
}

func similarity(list1, list2 []int) (int, error) {
	var similarity int

	for _, number := range list1 {
		count := 0

		for _, n := range list2 {
			if number == n {
				count++
			}
		}

		similarity += number * count
	}

	return similarity, nil
}

func main() {
	var list1, list2 []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	list1, list2, err = parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	d, err := distance(list1, list2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	s, err := similarity(list1, list2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Printf("distance: %d\nsimilarity: %d", d, s)
}

package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func partOne(fileLines []string) int64 {
	leftList := &IntHeap{}
	rightList := &IntHeap{}

	for i, line := range fileLines {
		data := strings.Split(line, "   ")

		leftData, err := strconv.ParseInt(data[0], 10, 64)
		if err != nil {
			fmt.Printf("Error reading left data on line %d", i)
		}

		leftList.Push(leftData)

		rightData, err := strconv.ParseInt(data[1], 10, 64)
		if err != nil {
			fmt.Printf("Error reading right data on line %d", i)
		}

		rightList.Push(rightData)
	}

	heap.Init(leftList)
	heap.Init(rightList)

	var sum int64 = 0
	for (leftList.Len() > 0) && (rightList.Len() > 0) {
		leftValue := heap.Pop(leftList).(int64)
		rightValue := heap.Pop(rightList).(int64)

		sum += int64(math.Abs(float64(leftValue - rightValue)))
	}

	return sum
}

func partTwo(fileLines []string) int64 {
	var leftList []int64
	rightMap := make(map[int64]int)

	for i, line := range fileLines {
		data := strings.Split(line, "   ")

		leftData, err := strconv.ParseInt(data[0], 10, 64)
		if err != nil {
			fmt.Printf("Error reading left data on line %d", i)
		}

		leftList = append(leftList, leftData)

		rightData, err := strconv.ParseInt(data[1], 10, 64)
		if err != nil {
			fmt.Printf("Error reading right data on line %d", i)
		}

		rightMap[rightData] += 1

	}

	var sum int64 = 0
	for _, num := range leftList {
		sum += int64(math.Abs(float64(num * int64(rightMap[num]))))
	}

	return sum
}

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error opening file")
	}

	fileLines := strings.Split(string(fileBytes), "\n")
	if len(fileLines) > 0 {
		fileLines = fileLines[:len(fileLines)-1]
	}

	start := time.Now()

	sum := partTwo(fileLines)

	elapsed := time.Since(start)
	fmt.Printf("took %s", elapsed)
	fmt.Printf("diff: %d\n", sum)
}

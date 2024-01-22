package main

import "fmt"

type SortAlgo interface {
	sort(list []int)
}

type BubbleSort struct{}

func (b *BubbleSort) sort(list []int) {
	n := len(list)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				// Swap if the element found is greater than the next element
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

type SelectionSort struct{}

func (b *SelectionSort) sort(list []int) {
	n := len(list)

	for i := 0; i < n-1; i++ {
		// Assume the current index is the minimum
		minIndex := i

		// Find the index of the minimum element in the unsorted part
		for j := i + 1; j < n; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		// Swap the found minimum element with the element at index i
		list[i], list[minIndex] = list[minIndex], list[i]
	}
}

type Context struct {
	SortAlgorithm SortAlgo
}

func (c *Context) SetSortAlgorithm(algo SortAlgo) {
	c.SortAlgorithm = algo
}

func (c *Context) ExecuteSort(list []int) {
	c.SortAlgorithm.sort(list)
}

func main() {
	context := &Context{
		SortAlgorithm: &BubbleSort{},
	}

	list := []int{64, 34, 25, 12, 22, 11, 90}

	context.ExecuteSort(list)
	fmt.Println("Sorted list using Bubble Sort:", list)

	list = []int{64, 34, 25, 12, 22, 11, 90}

	context.SetSortAlgorithm(&SelectionSort{})
	context.ExecuteSort(list)
	fmt.Println("Sorted list using Selection Sort:", list)
}

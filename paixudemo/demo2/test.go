package main

import "fmt"

func BubbleSort(selice *[]int) {
	for j := 0; j < len(*selice)-1; j++ {
		for i := 0; i < len(*selice)-1-j; i++ {
			if (*selice)[i] > (*selice)[i+1] { //大小
				(*selice)[i], (*selice)[i+1] = (*selice)[i+1], (*selice)[i]
			}
		}
	}
	fmt.Println("冒泡排序", *selice)
}

func SelectSort(selice *[]int) {

	for j := 0; j < len((*selice))-1; j++ {
		max := (*selice)[j]
		maxIndex := j
		for i := j + 1; i < len((*selice)); i++ {
			if max > (*selice)[i] { //大小
				max = (*selice)[i]
				maxIndex = i
			}
		}
		if maxIndex != j {
			(*selice)[j], (*selice)[maxIndex] = (*selice)[maxIndex], (*selice)[j]
		}
	}
	fmt.Println("选择排序", *selice)
}

func main() {

	selice := []int{6, 9, 2, 33, 56, 3, 20}
	//BubbleSort(&selice)
	SelectSort(&selice)

}

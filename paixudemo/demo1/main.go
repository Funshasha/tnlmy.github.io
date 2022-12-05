package main

import (
	"fmt"
)

//冒泡排序
func bubbleSort(selice *[]int) {

	for j := 0; j < len((*selice))-1; j++ {
		for i := 0; i < len((*selice))-1-j; i++ {
			if (*selice)[i] > (*selice)[i+1] {
				(*selice)[i], (*selice)[i+1] = (*selice)[i+1], (*selice)[i]
			}
		}
	}
	fmt.Println("冒泡排序法:", *selice)
}

//选择排序法
func selectSort(selice *[]int) {

	for j := 0; j < len((*selice))-1; j++ {
		max := (*selice)[j]
		maxIndex := j
		for i := j + 1; i < len((*selice)); i++ {
			if max < (*selice)[i] {
				max = (*selice)[i]
				maxIndex = i
			}
		}
		//交换
		if maxIndex != j {
			(*selice)[j], (*selice)[maxIndex] = (*selice)[maxIndex], (*selice)[j]
		}
		fmt.Printf("第%d次 %v\n", j+1, *selice)
	}
}

//插入排序

func insertSort(selice *[]int) {

	for i := 1; i < len((*selice)); i++ {
		insertVal := (*selice)[i]
		insertIndex := i - 1 //下标
		//从小到大
		for insertIndex >= 0 && (*selice)[insertIndex] < insertVal {
			(*selice)[insertIndex+1] = (*selice)[insertIndex] //数据后移
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			(*selice)[insertIndex+1] = insertVal
		}
		fmt.Printf("插入排序 第%d次插入后 %v\n", i, *selice)
	}

}

func quickSort(left int, right int, selice *[]int) {
	l := left
	r := right
	//pivot 是中轴，支点
	pivot := (*selice)[(left+right)/2]

	//for 循环的目标是将pivot 小的数放到左边
	//比pivot 大的数放在右边、
	for l < r {
		for (*selice)[l] < pivot {
			l++
		}
		for (*selice)[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		(*selice)[l], (*selice)[r] = (*selice)[r], (*selice)[l]
		if (*selice)[l] == pivot {
			r--
		}
		if (*selice)[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		quickSort(left, r, selice)
	}
	if right > l {
		quickSort(l, right, selice)
	}
}

func main() {

	selice := []int{52, 89, 0, 7, 66}
	// bubbleSort(&selice)
	// selectSort(&selice)
	// insertSort(&selice)

	quickSort(0, len(selice)-1, &selice)
	fmt.Println(selice)

	//BubbleSort(&selice)
}

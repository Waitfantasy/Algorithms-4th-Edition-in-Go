package main

import "fmt"

func sort(a []int) {

	// 利用自然的归并排序进行自底向上的排序。
	// 充分利用已经有序的数组, 找到第一个有序数组, 继续找到第二个有序数组, 然后将两个数组进行合并
	aux := make([]int, len(a))
	for {
		lo := 0

		// 找到第一个块, 设置mid的值为第一个块的index值
		mid := findArr(a, lo) - 1

		if mid == len(a)-1 {
			break
		}

		for mid < len(a)-1 {
			// 找到第二个块
			// 合并第一个与第二个块
			hi := findArr(a, mid+1) + mid
			merge(lo, mid, hi, a, aux)

			// 移动lo为第二个块的位置
			lo = hi + 1

			// 继续找第一个数组
			mid = findArr(a, lo) + lo
		}
	}
}

func findArr(a []int, lo int) int {
	index := 1
	for i := lo; i < len(a)-1; i++ {
		if a[i] <= a[i+1] {
			index++
		} else {
			break
		}
	}

	return index
}

func merge(lo, mid, hi int, a, aux []int) {
	i := 0
	j := mid + 1

	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[j] < aux[i] {
			a[k] = aux [j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func main() {
	a := []int{1, 1, 2, 3, 4, 2,}
	sort(a)
	fmt.Println(a)
}

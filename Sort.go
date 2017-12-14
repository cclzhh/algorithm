package main

import (
	"fmt"
	"sort"
)

func QuickSortM(a sort.Interface, start, end int) {
	// 退出条件
	if end - start < 1 { return }
	// 分区, 使用 end-1元素为pivot
	head, tail := start, end-1

	// divide and conquer
	for head < tail {
		for head < tail && a.Less(head, end) { head ++ }
		for head < tail && a.Less(end, tail) { tail -- }
		a.Swap(tail, head)
	}
	// 处理pivot
	if a.Less(end, head) {
		a.Swap(end, head)
	} else {
		head ++
	}
	// 递归调用
	QuickSortM(a, start, head-1)
	QuickSortM(a, head+1, end)
}

func BubbleSortM(a sort.Interface, start, end int) {
	for i := start; i < end; i++ {
		isChanged := false
		for j := start; j < end-i; j++ {
			if a.Less(j+1, j) {
				a.Swap(j, j+1)
				isChanged = true
			}
		}
		// 优化
		if isChanged == false {
			break
		}
	}
}

func InsertionSortM(a sort.Interface, start, end int) {
	for i := start+1; i < end; i++ {
		for j := i; j > start && a.Less(j, j-1) ; j-- {
			a.Swap(j, j-1)
		}
	}
}

func HeapSortM(a sort.Interface, start, end int) {
	// 初始化堆
	for i := end/2; i >= start; i-- {
		MaxHeapify(a, i, end)
	}
	// 排序
	for i := end; i > start; i-- {
		// 交换堆顶
		a.Swap(start, i)
		// 调整堆(剩余元素)
		MaxHeapify(a, start, i-1)
	}
}

// 大顶堆调整，堆顶元素可能不满足堆条件(因为排序过程中替换了堆顶元素，或者建堆时每次只增加一层)
func MaxHeapify(a sort.Interface, start, end int) {
	dad := start
	son := dad * 2 + 1
	for son <= end {
		if son + 1 <= end && a.Less(son, son+1) { son ++ }  // 找大的
		if a.Less(dad, son) {
			a.Swap(dad, son)
			dad = son
			son = dad * 2 + 1
		} else {  // 父节点已经最大则退出
			break
		}
	}
}

func main() {
	a := []int{1,3,2,5,-1,2,7,9,0,100}
	//QuickSortM(sort.IntSlice(a), 0, len(a)-1)
	//BubbleSortM(sort.IntSlice(a), 0, len(a)-1)
	//InsertionSortM(sort.IntSlice(a), 0, len(a)-1)
	HeapSortM(sort.IntSlice(a), 0, len(a)-1)

	fmt.Println(a)
}

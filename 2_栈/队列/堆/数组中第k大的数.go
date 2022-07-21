package main

/*
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
*/

//思路 第K大，取大小为K的最小堆
func findKthLargest(nums []int, k int) int {
	hp := &Heap{size: k}
	for _, num := range nums {
		hp.Add(num)
	}
	return hp.arr[0]
}

// 最小堆定义:除了叶子节点，其余父节点均大于他的子节点的完全二叉树
//难点，手写大根堆 or 小根堆
type Heap struct {
	arr  []int
	size int
}

func (hp *Heap) Add(num int) {
	if len(hp.arr) < hp.size {
		hp.arr = append(hp.arr, num)
		for i := len(hp.arr) - 1; i > 0; {
			p := (i - 1) / 2
			if p >= 0 && hp.arr[p] > hp.arr[i] {
				hp.Swap(p, i)
				i = p
			} else {
				break
			}
		}
	} else if num > hp.arr[0] {
		hp.arr[0] = num
		hp.Down(0)
	}
}

func (hp *Heap) Swap(a, b int) {
	hp.arr[a], hp.arr[b] = hp.arr[b], hp.arr[a]
}

func (hp *Heap) Down(i int) {
	k := i
	l, r := 2*i+1, 2*i+2
	n := len(hp.arr)
	if l < n && hp.arr[k] > hp.arr[l] {
		k = l
	}
	if r < n && hp.arr[k] > hp.arr[r] {
		k = r
	}
	if i != k {
		hp.Swap(i, k)
		hp.Down(k)
	}
}

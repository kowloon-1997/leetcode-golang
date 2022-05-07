package dhf

/*
递归:即在函数中调用自己的函数， 通过简短的两个函数即可实现复杂的计算逻辑

递归函数要素: 循环体 and 终止条件
*/

//eg 求1加到100
func digui(val, sum int64) int64 {
	if val == 100 {
		return sum
	}

	sum += val
	return digui(val+1, sum)
}

/*
回溯法

回溯法：又称为试探法，顾明思意就是每走一步就判断一下，如果合适就继续走，如果不合适就退一步再继续走
*/

/*
给定一个无重复的数组， 都组成所有无重复的子数组

构思：【1，2，3】

1  -> 12  -> 123
		  -> 12
   -> 1   -> 13
          -> 1
[] -> [2] -> [2,3]
	      -> [2]
   -> []  -> [3]
		  -> []


*/
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	temp := make([]int, 0)
	generate(0, nums, temp, &ret)
	return ret
}

func generate(i int, nums []int, temp []int, result *[][]int) {
	if i == len(nums)-1 {
		*result = append(*result, temp)
	} else {
		generate(i+1, nums, temp, result)
		temp = append(temp, nums[i])
		generate(i+1, nums, temp, result)
	}
}

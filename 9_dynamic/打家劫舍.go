package dynamic

/*
一条直线上有n个报警器， 每个房间中有若干不等金银财宝， 由于相邻的房间中安装有报警器， 偷取相邻房间的会触发报警， 那请问在不触发报警器的情况下，最多可以偷多少金银财宝

构思：
[1，2，4，6，3，5，6，2，5，7，4，3]

拆分问题:xi = max(i + xi-1)(xi-1)
确定状态:xi 表示第i个偷与不偷的最大值，
初始条件:x1 = 1, x2 = x2  x3=max(x1+3, x2)
转移方程:dp[i] = max(dp[i-1], i+dp[i-2])
coding:


*/

func Dajiajieshe(arr []int64) int64 {
	ret := make([]int64, len(arr))
	if len(arr) == 1 {
		return arr[0]
	}

	ret[0] = arr[0]
	if arr[0] > arr[1] {
		ret[1] = arr[0]
	} else {
		ret[1] = arr[1]
	}
	for i := 2; i < len(arr); i++ {
		if ret[i-2]+arr[i] > ret[i-1] {
			ret[i] = ret[i-2] + arr[i]
		} else {
			ret[i] = ret[i-1]
		}
	}
	return ret[len(arr)-1]
}

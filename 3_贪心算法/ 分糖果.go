package main

import "sort"

func main() {

}

//两个数组都排序后， 根据小孩的需求因子， 从小到大依次进行分配， 直到所有的孩子都分配完
//因为糖果也是从小往大的进行分配， 所以下一个满足的糖果一定是在上一个满足的右边
func findContentChildren(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool {
		return g[i] <= g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] <= s[j]
	})

	r := 0
	for _, gItem := range g {
		for idx, sItem := range s {
			if gItem <= sItem {
				r++
				if idx == 0 {
					s = s[idx+1 : len(s)]
					break
				}
				sTemp := make([]int, 0)
				for i := 0; i < idx; i++ {
					sTemp = append(sTemp, s[i])
				}
				for i := idx + 1; i < len(s); i++ {
					sTemp = append(sTemp, s[i])
				}
				s = sTemp
				break
			}
		}
	}
	return r
}

func findContentChildrenII(g []int, s []int) int {

	sort.Slice(g, func(i, j int) bool {
		return g[i] <= g[j]
	})

	sort.Slice(s, func(i, j int) bool {
		return s[i] <= s[j]
	})

	r := 0
	sIdx := 0

	for _, gItem := range g {
		for sIdx < len(s) {
			if s[sIdx] >= gItem {
				r++
				sIdx++
				break
			}
			sIdx++
		}
	}
	return r
}

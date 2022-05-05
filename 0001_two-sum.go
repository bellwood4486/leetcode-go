package leetcode_go

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		m[v] = i
	}
	for i1, v1 := range nums {
		v2 := target - v1
		if i2, ok := m[v2]; ok && i1 != i2 {
			return []int{i1, i2}
		}
	}
	return []int{}
}

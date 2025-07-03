package questions

func TwoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := dict[nums[i]]; ok {
			return []int{value, i}
		}
		dict[target-nums[i]] = i
	}
	return []int{0, 0}
}

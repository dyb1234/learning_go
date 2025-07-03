package questions

func RemoveDuplicates(nums []int) int {
	left := 0
	right := 0
	for right < len(nums) {
		if nums[right] == nums[left] {
		} else {
			left++
			nums[left] = nums[right]
		}
		right++
	}
	return left + 1
}

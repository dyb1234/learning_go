package questions

func PlusOne(digits []int) []int {
	digit := (digits[len(digits)-1] + 1) % 10
	carry := (digits[len(digits)-1] + 1) / 10
	digits[len(digits)-1] = digit
	for i := len(digits) - 2; i >= 0; i-- {
		digit = (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		digits[i] = digit
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

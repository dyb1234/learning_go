package questions

func IsPalindrome(x int) bool {
	y := 0
	z := x
	for x > 0 {
		y = y*10 + x%10
		x /= 10
	}
	return z == y
}

package questions

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; i < len(strs[0]); i++ {
		flag := false
		for _, str := range strs {
			if len(str) <= i {
				flag = true
				break
			}
			if str[i] != strs[0][i] {
				return prefix
			}
		}
		if !flag {
			prefix += string(strs[0][i])
		}
	}
	return prefix
}

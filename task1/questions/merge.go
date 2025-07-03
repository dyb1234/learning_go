package questions

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	temp := intervals[0]
	result = append(result, temp)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= temp[1] {
			if intervals[i][1] > temp[1] {
				temp[1] = intervals[i][1]
			}
			continue
		}
		temp = intervals[i]
		result = append(result, temp)
	}
	return result
}

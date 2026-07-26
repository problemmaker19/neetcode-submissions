func twoSum(nums []int, target int) []int {
    differences := make(map[int]int)
	result := []int{}

	for i, v := range nums {
		if val, ok := differences[target - v]; ok {
			result = append(result, val, i)
			break
		}
		differences[v] = i
	}

	return result
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)
	result := [][]string{}

	for _, word := range strs {
		frequency := [26]int{}
		for _, letter := range word {
			frequency[letter - 97] += 1
		}
		groups[frequency] = append(groups[frequency], word)
	}

	for _, g := range groups {
		result = append(result, g)
	}

	return result
}

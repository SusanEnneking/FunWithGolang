package twosum

func TwoSum(input []int, target int) []int {
	map1 := make(map[int]int)
	if len(input) < 2 {
		return []int{}
	}
	for i := 0; i < len(input); i++ {
		key := target - input[i]
		_, ok := map1[key]

		if ok {
			if key*2 == target {
				// special case where number is even and its halve occurs twice in the input
				return []int{map1[key], i}
			}
		} else {
			map1[key] = i
		}
	}
	// if there is an entry, there should be two entries
	for k := range map1 {
		index1 := map1[k]
		k2 := target - k
		index2, ok := map1[k2]
		//if they are equal that should have been found in the first loop
		if ok && index2 != index1 {
			return []int{index1, index2}
		}
	}
	return []int{}
}

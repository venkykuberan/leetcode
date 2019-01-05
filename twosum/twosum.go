package twosum

func twoSum(list []int, target int) []int {

	indxmap := make(map[int]int)
	for i := 0; i < len(list); i++ {
		comp := target - list[i]
		if _, ok := indxmap[comp]; ok {
			return []int{indxmap[comp], i}
		}
		indxmap[list[i]] = i
	}
	return nil

}

package compareTwoSlices

const DELETE = "delete"
const ADD = "add"
const SAME = "same"

func index(s []int64, v int64) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}
func remove(s []int64, value int64) []int64 {
	i := index(s, value)
	if i > -1 {
		s = append(s[:i], s[i+1:]...)
	}
	return s
}

func contains(s []int64, value int64) bool {
	return index(s, value) >= 0
}

func CompareTwoSlices(oldSlice []int64, newSlice []int64) map[string][]int64 {
	result := make(map[string][]int64)

	for _, item := range oldSlice {
		result[DELETE] = append(result[DELETE], item)
	}
	for _, item := range newSlice {
		if _, exists := result[DELETE]; exists {
			if contains(result[DELETE], item) {
				result[SAME] = append(result[SAME], item)
				result[DELETE] = remove(result[DELETE], item)
			} else {
				result[ADD] = append(result[ADD], item)
			}
		} else {
			result[ADD] = append(result[ADD], item)
		}
	}
	return result
}

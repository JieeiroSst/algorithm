package Linear_Search

func LinearSearch(datalist []int, key int) bool{
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

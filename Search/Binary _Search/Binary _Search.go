package Binary__Search

func BinarySearch(needle int, haystack []int) bool {
	low:=0
	hight:=len(haystack) -1

	for low <= hight {
		midian:= (low+hight)/2

		if haystack[midian] < needle {
			low = midian + 1
		}else {
			hight = midian - 1
		}
	}

	if low == len(haystack) || haystack[low] !=needle {
		return false
	}

	return true
}
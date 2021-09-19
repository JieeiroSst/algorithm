package Interpolation__Search

func InterpolationSearch(array []int, key int) int {
	min , max:= array[0] , array[len(array) -1]

	low , hight:= 0 , len(array) -1

	for {
		if key < min {
			return low
		}

		if key > max {
			return hight + 1
		}

		var guess int

		if hight == low {
			guess = hight
		}else {
			size:= hight - low
			offset:=int(float64(size-1)*(float64(key-min)/float64(max-min)))
			guess = low+ offset
		}

		if array[guess] == key {
			for guess > 0 && array[guess-1] == key {
				guess--
			}
			return guess
		}

		if array[guess] > key {
			hight = guess - 1
			max = array[hight]
		}else{
			low = guess + 1
			min = array[low]
		}
	}

}

package Bubble_Sort

func BubbleSort(items []int) {
	var (
		n = len(items)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}

func bubbleSort(arr[]int){
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1]= arr[j+1],arr[j]
			}
		}
	}
}
package util

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[i] > arr[i+1] {
				temp := arr[i+1]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	return arr
}

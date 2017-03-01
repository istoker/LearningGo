package selsort

//indexOfMinimal returns the index of the smallest integer from a slice
func indexOfMinimal(arr []int) int  {
    minValue := arr[0]
    minIndex := 0
    for i, v := range arr[minIndex:]{
        if v < minValue{
            minValue = v
            minIndex = i
        }
    }
    return minIndex
}

//swap swaps the elements of arr at index1 and index2
func swap(arr []int, index1 int, index2 int) []int{
    arr[index1], arr[index2] = arr[index2], arr[index1]
    return arr
}

//SelectionSort Returns the slice sorted with selection sort
func SelectionSort(arr []int) []int{
    for i := range arr{
        arr = swap(arr, i, i+indexOfMinimal(arr[i:]))
    }
    return arr
}




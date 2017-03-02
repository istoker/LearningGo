package mergesort

//Merges the two sorted subslices, which are indicated by half.
func merge(arr []int, half int){
    lowSlice := make([]int, len(arr[:half]))
    copy(lowSlice, arr[:half])
    highSlice := make([]int, len(arr[half:]))
    copy(highSlice, arr[half:])
    i, j, k := 0, 0, 0;
    for i < len(lowSlice) && j < len(highSlice){
        if lowSlice[i] <= highSlice[j]{
            arr[k] = lowSlice[i]
            i++
        } else{
            arr[k] = highSlice[j]
            j++
        }
        k++
    }
    if i < len(lowSlice){
        for i < len(lowSlice){
            arr[k] = lowSlice[i]
            i++
            k++
        }
    } else{
        for j < len(highSlice){
            arr[k] = highSlice[j]
            j++
            k++
        }
    }
}

//MergeSort sorts the array using the mergesort algortihm
func MergeSort(arr []int){
    if(len(arr) > 1){
        half := len(arr)/2
        MergeSort(arr[:half])
        MergeSort(arr[half:])
        merge(arr, half)
    }
}
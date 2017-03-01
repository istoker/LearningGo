package inssort

//Inserts the last element of the slice into the correct place within the slice.
//Will slide the elements bigger then the last element to the right.
func insert(arr []int){
    key := arr[len(arr) - 1]
    inserted := false
    for i := len(arr) - 2; i >= 0; i--{
        if arr[i] > key{
            arr[i+1] = arr[i]
        }else{
            arr[i+1] = key
            inserted = true
        }
    }
    if !inserted{
        arr[0] = key
    }
}

//InsertionSort sorts the underlying array of the slice via the insertion sort algorithm
func InsertionSort(arr []int){
    for i := 0; i < len(arr); i++{
        insert(arr[:i+1])
    }
}
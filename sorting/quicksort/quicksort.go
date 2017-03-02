package quicksort

func swap(arr []int, index1 int, index2 int){
    arr[index1], arr[index2] = arr[index2], arr[index1]
}

func partition(arr []int) int{
    r := len(arr)-1 //r = pivot
    j, q := 0, 0 //j = current index of element to check, q = first index of element > than pivot
    for j < len(arr)-1{
        if arr[j] > arr[r]{
            j++
        }else{
            swap(arr, j, q)
            j++ 
            q++
        }
    }
    swap(arr, r, q)
    return q
}

//Quicksort sorts the slice using the quicksort algorithm
func Quicksort(arr []int){
    if(len(arr) > 1){
        q := partition(arr)
        Quicksort(arr[:q])
        Quicksort(arr[q+1:])
    }
}
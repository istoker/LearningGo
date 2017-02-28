package binsearch

//BinarySearch searches for target within arr using the binary search algorithm. arr must be sorted.
func BinarySearch(arr []int, target int) (int, int){
    min := 0
    max := len(arr)-1
    tries := 0
    for max >= min{
        var guess = (min+max)/2
        tries++
        if arr[guess] == target{
            return guess, tries
        } else if arr[guess] < target{
            min = guess + 1
        } else{
            max = guess - 1;
        }
    }
    return -1, tries
}
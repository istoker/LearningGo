package main

import "fmt"
import "github.com/istoker/LearningGO/sorting/selsort"

func main(){
    arr := make([]int, 20)
    for i := range arr{
        arr[i] = 19 - i
    }
    fmt.Println("Original slice:")
    fmt.Println(arr)
    selSortArr := selsort.SelectionSort(arr[:])
    fmt.Println("slice sorted with selection sort")
    fmt.Println(selSortArr)
}
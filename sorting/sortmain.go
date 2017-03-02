package main

import "fmt"
import "github.com/istoker/LearningGO/sorting/selsort"
import "github.com/istoker/LearningGo/sorting/inssort"
import "github.com/istoker/LearningGO/sorting/mergesort"

func main(){
    arr := [20]int{19,18,17,16,14,15,13,12,10,9,11,8,7,4,5,6,3,2,1,0}
    fmt.Println("Original Array:")
    fmt.Println(arr)
    selSortArr := arr
    insSortArr := arr
    merSortArr := arr
    selsort.SelectionSort(selSortArr[:])
    inssort.InsertionSort(insSortArr[:])
    mergesort.MergeSort(merSortArr[:])
    fmt.Println("Array sorted with selection sort:")
    fmt.Println(selSortArr)
    fmt.Println("Array sorted with insertion sort:")
    fmt.Println(insSortArr)
    fmt.Println("Array sorted with merge sort:")
    fmt.Println(merSortArr)
}
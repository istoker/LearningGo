package main

import "fmt"
import "github.com/istoker/LearningGO/sorting/selsort"
import "github.com/istoker/LearningGo/sorting/inssort"
import "github.com/istoker/LearningGO/sorting/mergesort"
import "github.com/istoker/LearningGO/sorting/quicksort"
import "time"
import "math/rand"

//from: https://coderwall.com/p/cp5fya/measuring-execution-time-in-go
func timeTrack(start time.Time, name string){
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}

func selectionSort(arr []int){
    defer timeTrack(time.Now(), "Selection sort")
    selsort.SelectionSort(arr)
}

func insertionSort(arr []int){
    defer timeTrack(time.Now(), "Insertion sort")
    inssort.InsertionSort(arr)
}

func mergeSort(arr []int){
    defer timeTrack(time.Now(), "Merge sort")
    mergesort.MergeSort(arr)
}

func quickSort(arr []int){
    defer timeTrack(time.Now(), "Quick sort")
    quicksort.Quicksort(arr)
}

func main(){
    arr := [50000]int{}
    rand.Seed(time.Now().Unix())
    for i:= 0; i < len(arr); i++{
        arr[i] = rand.Intn(len(arr))
    }
    fmt.Println("Original Array:")
    fmt.Println(arr[:10])
    selSortArr := arr
    insSortArr := arr
    merSortArr := arr
    quiSortArr := arr
    
    selectionSort(selSortArr[:])
    fmt.Println("Array sorted with selection sort:")
    fmt.Println(selSortArr[:10])
    
    insertionSort(insSortArr[:])
    fmt.Println("Array sorted with insertion sort:")
    fmt.Println(insSortArr[:10])
    
    mergeSort(merSortArr[:])
    fmt.Println("Array sorted with merge sort:")
    fmt.Println(merSortArr[:10])
    
    quickSort(quiSortArr[:])
    fmt.Println("Array sorted with quick sort:")
    fmt.Println(quiSortArr[:10])
}
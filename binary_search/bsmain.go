package main

import "fmt"
import "github.com/istoker/LearningGo/binary_search/binsearch"
import "math/rand"
import "time"

func init(){
    rand.Seed(time.Now().Unix())
}

func main(){
    primes := [25]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
    var targetIndex = rand.Intn(25)
    var target = primes[targetIndex]
    foundIndex, tries := binsearch.BinarySearch(primes[:], target)
    fmt.Printf("Target index is: %d\n", targetIndex)
    fmt.Printf("Found index is: %d\n", foundIndex)
    if foundIndex == targetIndex {
        fmt.Printf("Success in %d tries\n", tries)
    } else{
        fmt.Printf("Failure in %d tries\n", tries)
    }
}


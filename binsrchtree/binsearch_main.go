package main

import "fmt"
import "github.com/istoker/LearningGo/binsrchtree/bst"

func main(){
	boom := new(bst.Tree)
	boom.Insert(50, "root")
	boom.Insert(20, "leftchild")
	boom.Insert(30, "rightchild")
	boom.Insert(10, "leftchild")
	boom.Insert(05, "leftchild")
	boom.Insert(25, "leftchild")
	boom.Insert(35, "rightchild")
	boom.Insert(75, "rightchild")
	fmt.Println(boom)
	fmt.Printf("Size: %d\n", boom.Size())
	boom.PrintTree()
	boom.Delete(30)
	boom.PrintTree()

}
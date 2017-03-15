package bst

import "fmt"

//Node is a node for the binary search tree
type Node struct{
	key 	int
	data	string
	left	*Node
	right	*Node
}

//Tree is an unbalanced binary search tree
type Tree struct{
	root *Node
}

func initializeNode(k int, d string) *Node{
	return &Node{key : k, data : d}
}

func subTreeSize(n *Node) int{
	if n.left != nil && n.right != nil{
		return subTreeSize(n.left) + 1 + subTreeSize(n.right)
	}else if n.left != nil{
		return subTreeSize(n.left) + 1
	}else if n.right != nil{
		return subTreeSize(n.right) + 1
	}
	return 1

}

func findKey(n *Node, key int) (string, bool){
	if n.key == key{
		return n.data, true
	}
	if key < n.key{
		if n.left != nil{
			return findKey(n.left, key)
		}
	}else{
		if n.right != nil{
			return findKey(n.right, key)
		}
	}
	return "", false
}

func insertNode(n *Node, k int, d string) bool{
	if k == n.key{
		return false
	}
	if k < n.key{
		if n.left != nil{
			return insertNode(n.left, k, d)
		}
		n.left = initializeNode(k, d)
	} else{
		if n.right != nil{
			return insertNode(n.right, k, d)
		}
		n.right = initializeNode(k, d)
	}
	return true
}

func findMax(n *Node) *Node{
	if n.right != nil{
		return findMax(n.right)
	}
	return n
}

func findMin(n *Node) *Node{
	if n.left != nil{
		return findMin(n.left)
	}
	return n
}

func deleteFromParent(n *Node, par *Node) bool{
	fmt.Printf("Deleting %d\n", n.key)
	if n.left != nil && n.right != nil{
		max := findMax(n.left)
		temp := *max
		deleteNode(n, max.key)
		temp.left = n.left
		temp.right = n.right
		*n = temp
	}else if n.left != nil{
		if n.key < par.key{
			par.left = n.left
		}else{
			par.right = n.left
		}
	}else if n.right != nil{
		if n.key < par.key{
			par.left = n.right
		}else{
			par.right = n.right
		}
	}else{
		if n.key < par.key{
			par.left = nil
		}else{
			par.right = nil
		}
	}
	return true
}

//Gaat ervan uit dat de node altijd bestaat. Dit moet gefixed worden.
func deleteNode(n *Node, key int) bool{
	if key < n.key{
		if n.left != nil{
			if n.left.key == key{
				return deleteFromParent(n.left, n)
			}
			return deleteNode(n.left, key)
		}
	}else{
		if n.right != nil{
			if n.right.key == key{
				return deleteFromParent(n.right, n)
			}
			return deleteNode(n.right, key)
		}
	}
	return false
}

func traverseNodes(n *Node) []*Node{
	var l []*Node
	var r []*Node
	if n.left != nil{
		l = traverseNodes(n.left)
	}
	if n.right != nil{
		r = traverseNodes(n.right)
	}
	s := append(l, n)
	s = append(s, r...)
	return s
}

//===== Tree functions =====

//Size returns the number of nodes in the Tree
func (t *Tree) Size() int{
	if t.root != nil{
		return subTreeSize(t.root)
	}
	return 0
}

//Delete deletes the node with key from the tree. returns false if node not found.
func (t *Tree) Delete(key int) bool{
	if t.root != nil{
		if t.root.key == key{
			t.root = nil
		}
		return deleteNode(t.root, key)
	}
	return false
}

//Insert inserts a new node with the corresponding data to the Tree
func (t *Tree) Insert(key int, data string) bool{
	if t.root != nil {
		return insertNode(t.root, key, data)
	}
	t.root = initializeNode(key, data)
	return true
}

//Search searches for the node with Key and returns its data and whether the node has been found.
func (t *Tree) Search(key int) (string, bool){
	if t.root != nil{
		return findKey(t.root, key)
	}
	return "", false
}

//Traverse returns a slice with pointers to the nodes in-order
func (t *Tree) Traverse() []*Node{
	if t.root != nil{
		return traverseNodes(t.root)
	}
	return make([]*Node, 0, 0)
}

//PrintTree prints the tree in order
func (t *Tree) PrintTree(){
	for _, v := range t.Traverse(){
		fmt.Printf("%d ", v.key)
	}
	fmt.Println()
}
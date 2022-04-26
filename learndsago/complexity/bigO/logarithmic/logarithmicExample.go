package logarithmic

import "fmt"

type Tree struct {
	LeftNode  *Tree
	Value     int
	RightNode *Tree
}

func (tree *Tree) insert(m int) {
	if tree != nil {
		if tree.LeftNode == nil {
			tree.LeftNode = &Tree{nil, m, nil}
		} else {
			if tree.RightNode == nil {
				tree.RightNode = &Tree{nil, m, nil}
			} else {
				if tree.LeftNode != nil {
					tree.LeftNode.insert(m)
				} else {
					tree.RightNode.insert(m)
				}
			}
		}
	} else {
		tree = &Tree{nil, m, nil}
	}
}

func print(tree *Tree) {
	if tree != nil {
		fmt.Println("Value:", tree.Value)
		fmt.Printf("Node Left:")
		print(tree.LeftNode)
		fmt.Printf("Node Right:")
		print(tree.RightNode)

	} else {
		fmt.Printf("null tree.")
	}
}

func Example() {
	var t *Tree = &Tree{nil, 1, nil}
	t.insert(3)
	t.LeftNode.insert(5)
	t.LeftNode.LeftNode.RightNode.insert(7)
	print(t)
}

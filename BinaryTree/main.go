package main

import (
	"fmt"
)

type Node struct {
	value     int
	RightNode *Node
	LeftNode  *Node
}

func (n *Node) Insert(value int) {
	root := n

	for true {
		if root.value < value {
			if root.RightNode == nil {
				setNode := Node{value: value, RightNode: nil, LeftNode: nil}
				root.RightNode = &setNode
				break
			}
			root = root.RightNode
		}
		if root.value > value {
			if root.LeftNode == nil {
				setNode := Node{value: value, RightNode: nil, LeftNode: nil}
				root.LeftNode = &setNode
				break
			}
			root = root.LeftNode
		}
	}
}
func (n *Node) Find(value int) bool {
	root := n
	for true {
		if value > root.value {
			if root.RightNode == nil {
				return false
			}
			root = root.RightNode
		}
		if value < root.value {
			if root.LeftNode == nil {
				return false
			}
			root = root.LeftNode
		}
		if value == root.value {
			return true
		}
	}
	return false
}
func Pre_Order_Travers(value Node) {

	if value.LeftNode != nil {
		Pre_Order_Travers(*value.LeftNode)
	}
	if value.RightNode != nil {
		Pre_Order_Travers(*value.RightNode)
	}
}

func In_Order_Travers(value Node) {
	root := value

	if root.LeftNode != nil {
		In_Order_Travers(*value.LeftNode)
	}
	fmt.Println(root)
	if root.RightNode != nil {
		In_Order_Travers(*value.RightNode)
	}
}
func Post_Order_Travers(value Node) {
	root := value

	if root.LeftNode != nil {
		Post_Order_Travers(*value.LeftNode)
	}
	fmt.Println(root)
	if root.RightNode != nil {
		Post_Order_Travers(*value.RightNode)
	}
}
func Height(value *Node) int {
	if value.RightNode == nil && value.LeftNode == nil {
		return 0
	}
	if value.RightNode == nil {
		return 1 + Height(value.LeftNode)
	}
	if value.LeftNode == nil {
		return 1 + Height(value.RightNode)
	}
	return 1 + max(Height(value.RightNode), Height(value.LeftNode))

}
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func Min_object(value *Node) int {
	if value.RightNode == nil && value.LeftNode == nil {
		return value.value
	} else if value.RightNode == nil {
		return value.LeftNode.value
	} else if value.LeftNode == nil {
		return value.RightNode.value
	}
	return min(Min_object(value.LeftNode), Min_object(value.RightNode))

}
func equals(value *Node, n *Node) bool {
	if n == nil && value == nil {
		return true
	}
	if n != nil && value != nil {
		return n.value == value.value &&
			equals(value.LeftNode, n.LeftNode) && equals(value.RightNode, n.RightNode)
	}
	return false
}

func main() {
	
	tree := Node{5, nil, nil}
	tree.Insert(1)
	
}

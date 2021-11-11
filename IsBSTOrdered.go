/*******************************************************************************************
Problem 1. Write the function "IsBSTOrdered()": If the values in the binary
tree rooted at `root` satisfy the binary search tree ordering, return true. If
the values do NOT all satisfy the binary search tree ordering, return false.
You may assume there are no duplicate items in the tree.
*******************************************************************************************/

// TreeNode is used in this and the next problem. It represents a node in a binary tree.
type TreeNode struct {
	left  *TreeNode // subtree with smaller values (nil if no subtree)
	right *TreeNode // subtree with larger values (nil if no subtree)
	value int
}

func IsBSTOrdered(root *TreeNode) bool {
	max_int := 999999
	min_int := -999999
	return IsBSTOrderedMinMax(root, min_int, max_int)
}

func IsBSTOrderedMinMax(root *TreeNode, min, max int) bool{
	if root == nil {
		return true
	}

	node_value := root.value

	if node_value > max {
		return false
	}

	if node_value < min {
		return false
	}

	return IsBSTOrderedMinMax(root.left, min, node_value) && IsBSTOrderedMinMax(root.right, node_value, max)
}

func InputQ1(){
	var root,c1,c2,c3,c4,c5,c6 TreeNode

	root.left = &c1
	root.right =&c2
	root.value = 4
	c1.left = &c3
	c1.right = &c4
	c2.left = &c5
	c2.right = &c6
	c3.left = nil
	c3.right = nil
	c4.left = nil
	c4.right = nil
	c5.left = nil
	c5.right = nil
	c6.left = nil
	c6.right = nil
	c1.value = 2
	c2.value = 15
	c3.value = 1
	c4.value = 3
	c5.value = 0
	c6.value = 120

	fmt.Println(IsBSTOrdered(&root))
}

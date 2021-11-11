/*******************************************************************************************
Problem 2. Lowest common ancestor. Given x and y, return the value in the tree
rooted at `root` that is in the node that is the lowest common ancestor of x
and y. The ancestors of a node x are all the nodes between x and the root. The
lowest common ancestor of two nodes x and y is the node that is an ancestor of
both x and y that is closest to x and y. You may assume nodes for x and y exist
in the tree, and that there are no nodes that share the same value in the tree.

Hint: Think about the BST tree ordering.
*******************************************************************************************/

func LowestCommonAncestor(root *TreeNode, x, y int) int {
	if root == nil {
		return 999999
	} else if root.value == x || root.value == y {
		return root.value
	}

	left_side :=  LowestCommonAncestor(root.left, x, y)
	right_side :=  LowestCommonAncestor(root.right, x, y)
	
	if left_side != 999999 {
		if right_side != 999999 {
			return root.value
		}
		return left_side
	}
	return right_side
}

func InputQ2(){
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
	c5.value = 10
	c6.value = 120

	fmt.Println(LowestCommonAncestor(&root,1,15))
}

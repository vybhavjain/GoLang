/*******************************************************************************************
Problem 7. Consider the following Node type. It contains a single field which
is a pointer to another Node. That other Node of course will itself contain a
field pointing to a Node.  By following these pointers we can walk through a
series of Nodes. If `next` is nil, the series of Nodes ends.

Write a function HasCycle that walks through the series of Nodes staarting with
the parameter n and returns true if you can ever return back to a node you have
already visited, and false otherwise.

Be sure to consider the following case:

n -> n1 -> n2 -> n3 -> n4 -> n5 -+
                 ^               |
                 |               |
                 +---------------+

which should return true.

You cannot modify the Node type or the pointers in the Node.
*******************************************************************************************/

type Node struct {
	next *Node
}

func HasCycle(n *Node) bool {

	visited_nodes := []Node{}

	for n.next!= nil {
		visited_address := *n.next
		for x := 0 ; x<len(visited_nodes); x++ {
			if visited_nodes[x] == visited_address {
				return true
			}
		}
		visited_nodes = append(visited_nodes,*n.next)
		n = n.next
		

	}
	return false
}


func main123() {
	var n,n1,n2,n3,n4,n5  Node 

	n = Node{&n1}
	n1 = Node{&n2}
	n2 = Node{&n3}
	n3 = Node{&n4}
	n4 = Node{&n5}
	n5 = Node{nil}

	fmt.Println(HasCycle(n.next))
}

func main() {

	k := [][]int{
		{3, 4, -5, 1, 7},
		{10, 10, 10},
		{2, 1, 0, -1},
		{3, -3, -3, 0},
		{4, -4, 5, 0},
	}

	fmt.Println(LeftmostNegative(k))
}
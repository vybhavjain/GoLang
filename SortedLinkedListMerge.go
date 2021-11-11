/*******************************************************************************************
Problem 3. Sorted linked list merge. You are given two linked lists, l1 and l2,
each of which contain values that are sorted by non-decreasing values. Return a
new linked list that contains all of the nodes of l1 and l2 and that is sorted.
You should not allocate new nodes; reuse the nodes in the input lists.

That is if l1 contains n items and l2 contains m items, the list you return
will have n+m items in sorted order.
*******************************************************************************************/

// ListNode is used in this problem and the next one. It is a node in a linked list.
type ListNode struct {
	value int
	next  *ListNode // next node, nil if at end.
}

func MergeSortedLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	} else if l2 == nil {
		return l1
	}

	if l1.value <= l2.value {
		l1.next = MergeSortedLists(l1.next, l2)
		return l1
	}else if l1.value > l2.value {
		l2.next = MergeSortedLists(l1, l2.next)
		return l2
	}
	return l1
}
//Non decreasing - increasing

func InputQ3(){
	var n1,n2,n3,m1,m2,m3 ListNode
	n1.value = 5
	n1.next = &n2
	n2.value = 8
	n2.next = &n3
	n3.value = 10
	n3.next = nil
	m1.value = 1
	m1.next = &m2
	m2.value = 8
	m2.next = &m3
	m3.value = 9
	m3.next = nil

	list := MergeSortedLists(&n1,&m1)
	for list!= nil {
		fmt.Println(list.value)
		list = list.next
	}
	
}

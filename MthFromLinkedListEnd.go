/*******************************************************************************************
Problem 4. Given an UNSORTED linked list, and an integer m, return the integer
(value) that is at position m from the END of the linked list. That is if m ==
0, return the last item in the linked list; if m == 1, return the second to
last, etc. You can assume that m is less than the length of the linked list.
*******************************************************************************************/

func MthFromEnd(ll *ListNode, m int) int {
	
	main_pointer := ll
	ref_pointer := ll

	for count:=0;count<m;count++ {
		main_pointer = main_pointer.next
	}  

	for main_pointer.next!= nil {
		main_pointer = main_pointer.next
		ref_pointer = ref_pointer.next
	}

	return ref_pointer.value 
}

func InputQ4(){
	var n1,n2,n3,n4,n5,n6,n7 ListNode
	var m int

	n1.value = 1
	n1.next = &n2
	n2.value = 6
	n2.next = &n3
	n3.value = 2
	n3.next = &n4
	n4.value = 8
	n4.next = &n5
	n5.value = 12
	n5.next = &n6
	n6.value = 3
	n6.next = &n7
	n7.value = 99
	n7.next = nil
	m = 3 
	fmt.Println(MthFromEnd(&n1,m))	
}

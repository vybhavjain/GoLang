/*******************************************************************************************
Problem 5. Stack with Min(). Modify the following Stack data structure to
support the following operations:

	Push(): push an item on the stack
	Pop(): pop an item from the stack
	Min(): return the value of the smallest integer on the stack (WITHOUT modifying the
		   items on the stack)

Each of the above operations must run in time *independent* of the number of
items on the stack. That is they should take O(1) time. Consequently, you
should not create any maps or arrays. You will have to modify the StackItem
and/or the Stack types and the Push and Pop operations and implement the Min
operation.

Note: you can create a new Stack with "var S = Stack{}" (since everything
defaults to nil); keep this approach to creating a new stack.

Example: The following code should work:

		var S = Stack{}
		S.Push(10)
		S.Push(3)
		S.Push(12)
		S.Push(2)
		fmt.Println(S.Min())	// prints 2
		S.Pop()
		fmt.Println(S.Min())	// prints 3
		S.Pop()
		fmt.Println(S.Min())	// prints 3

Hint: as one Push()es new items, the min can only go down.
*******************************************************************************************/
type Stack struct {
	top *StackItem // pointer to the item on the top of the stack
}

type StackItem struct {
	prev  *StackItem // pointer to the next item on the stack
	value int
	prev_ref *StackItem
	val_ref int
}

// Push adds a new item to the top of the stack. It runs in time independent
// of the number of elements in the stack.
func (s *Stack) Push(v int) {
	if s.top == nil {
		s.top = &StackItem{
		prev:  s.top,
		value: v,
		prev_ref:s.top ,
		val_ref:v , 
		}
	} else if s.top.val_ref < v  {
		s.top = &StackItem{
		prev:  s.top,
		value: v,
		prev_ref:s.top.prev_ref ,
		val_ref:s.top.val_ref , 
		}
	} else {
		s.top = &StackItem{
		prev:  s.top,
		value: v,
		prev_ref:s.top ,
		val_ref:v ,
		}
	}
}

// Pop removes and returns the item at the top of the Stack. It runs in
// time independent of the number of elements in the Stack.
func (s *Stack) Pop() int {
	if s.top == nil {
		panic("Pop on an empty stack!")
	}
	v := s.top.value
	s.top = s.top.prev
	return v
}

// Min returns the smallest integer on the stack without changing the items on the stack.
// It runs in time independent of the number of items in the Stack.
func (s *Stack) Min() int {
	return s.top.val_ref
}




func InputQ5(){
	var S = Stack{}
	S.Push(10)
	S.Push(3)
	S.Push(12)
	S.Push(2)
	fmt.Println("min is",S.Min())	// prints 2
	//fmt.Println("cycle 1 done")	// prints 2
	S.Pop()
	//fmt.Println("cycle 2 started")	// prints 2
	fmt.Println(S.Min())	// prints 3
	S.Pop()
	fmt.Println(S.Min())	// prints 3
}

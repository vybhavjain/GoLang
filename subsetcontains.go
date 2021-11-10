/*******************************************************************************************
Problem 5. Write a function Contains(l1, l2 []int) bool that returns true if
the integers in l2 are a subset  of the integers in l1 and false otherwise.
That is, the function returns true if every integer that appears in l2 also
appears someplace in l1. If an integer occurs in l2 more than once, it needs to
occur at least that many times in l1.
*******************************************************************************************/

func Contains(l1, l2 []int) bool {
	count := 0

	for x := 0;x<len(l2);x++ {
		for y := 0;y<len(l1);y++ {
			if l2[x] == l1[y] {
				l1 =append(l1[:y], l1[y+1:]...)
				count ++
				break
			}
		}
	}

	if count == len(l2) {
		return true	
	}
	
	return false
}

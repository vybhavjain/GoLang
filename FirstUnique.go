/*******************************************************************************************
Problem 3. Write a function called FirstUnique that takes a slice of integers
and returns the first integer that only occurs once in the slice. If no
integers occur only once, return 0.

Examples:
    FirstUnique([]int{2, 3, 4, 5, 2, 4, 5}) should return 3.
    FirstUnique([]int{7, 8, 2, 8, 2, 7, 8}) should return 0.
    FirstUnique([]int{8, 8, 8, 9, 8, 6, 8, 8}) should return 9.
*******************************************************************************************/
func FirstUnique(list []int) int {
    flag := 0
    element := 0

    for x:=0; x<len(list);x++ {
        flag = 0
        element = list[x]
        for y:=0; y<len(list);y++ {
            if x != y {
                if element == list[y] {
                    flag = 1
                }
            }
        }

        if flag == 0 {
            return element
        }
    }
    
    return 0
}

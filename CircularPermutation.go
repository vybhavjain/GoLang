/*******************************************************************************************
Problem 4. A list of numbers A is a circular permutation of another list B if
you can write both lists around a circle and rotate the circles so that the
positions of the numbers are identical.

For example:
    A = 1,7,8,10,31,14
    B = 10,31,14,1,7,8
    C = 7,8,1,10,14,31
A and B are circular permutations of each other, while C is not a circular
permutation of A or B.

Write a function IsCircularPermutation that takes two int lists as parameters
and returns true if one is a circular permutation of the other (and false
otherwise).
*******************************************************************************************/

func IsCircularPermutation(a, b []int) bool {
    
    len1:=len(a)
    len2:=len(b)

    if len1!=len2 {
        return false
    }

    for x:=0; x<len1;x++ {
        for y:=0; y<len2;y++ {
            count := 0
            if a[x] == b[y] {
                flag := 0
                i := x
                j := y
                for count < len1 {
                    i = i % len1
                    j = j % len2
                    if a[i] != b[j] {
                        flag = 1
                    }
                    i++
                    j++
                    count++
                } 
                if flag == 0 {
                    return true
                }
            }
        }
    }

    return false
}

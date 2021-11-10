func LeftmostNegative(f [][]int) int {
	var minimum int
	minimum = 999
	for x := 0; x<len(f);x++ {
		for y := 0;y<len(f[x]);y++ {
			element := f[x][y]
		
			if element < 0 {
				if y < minimum {
					minimum = y
				}
			}
		}
	}
	
	if minimum != 999 {
		return minimum
	}
	return -1
}

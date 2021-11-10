func CountSwitches(b []bool) int {
	a := 0
	for x := 0; x<len(b)-1;x++ {
		fmt.Println(x)
		if b[x] != b[x+1] {
			a += 1
		}
	}
	return a
}

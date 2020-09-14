package main

/**********************************************************************************************************************/
func SequenceSearch(a []int, b int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return i
		}
	}
	return -1
}

/**********************************************************************************************************************/
func BinarySearch0(a []int, b int) int {
	l, m, h := 0, 0, len(a)-1
	for l <= h {
		m = (h + l) / 2
		if a[m] == b {
			return m
		}
		if a[m] > b {
			h = m - 1
		}
		if a[m] < b {
			l = m + 1
		}
	}
	return -1
}

func BinarySearch1(a []int, b int) int {
	return bSearch(a, 0, len(a)-1, b)
}

func bSearch(a []int, l, h, b int) int {
	m := (l + h) / 2
	if a[m] == b {
		return m
	}
	if a[m] > b {
		return bSearch(a, l, m-1, b)
	}
	if a[m] < b {
		return bSearch(a, m+1, h, b)
	}
	return -1
}

/**********************************************************************************************************************/
func InsertSearch(a []int, b int) int {
	l, m, h := 0, 0, len(a)-1
	for l <= h {
		m = l + (b-a[l])/(a[h]-a[l])*(h-l)
		if a[m] == b {
			return m
		}
		if a[m] > b {
			h = m - 1
		}
		if a[m] < b {
			l = m + 1
		}
	}
	return -1
}

/**********************************************************************************************************************/
func FibonacciSearch(a []int, b int) int {
	var f []int
	f = append(f, 1, 1)
	for i := 2; i < 100; i++ {
		f = append(f, f[i-1]+f[i-2])
	}

	var n int
	for f[n]-1 < len(a) {
		n++
	}

	temp := make([]int, f[n]-1)
	for i, v := range a {
		temp[i] = v
	}
	for i := len(a); i < len(temp); i++ {
		temp[i] = a[len(a)-1]
	}

	l, m, h := 0, 0, len(a)
	for l <= h {
		m = l + f[n-1] - 1
		if temp[m] == b {
			if m < len(a) {
				return m
			} else {
				return len(a) - 1
			}
		}
		if temp[m] > b {
			h = m - 1
			n = n - 1
		}
		if temp[m] < b {
			l = m + 1
			n = n - 2
		}
	}

	return -1
}

/**********************************************************************************************************************/
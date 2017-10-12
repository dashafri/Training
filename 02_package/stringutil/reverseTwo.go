package stringutil

func reverseTwo(d string) string{
	x := []rune(d)
	for a, b := 0, len(x)-1; a < len(x)/2; a, b = a+1, b-1 {
		x[a], x[b] = x[b], x[a]
	}
	return string(x)
}
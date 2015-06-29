package simple

func Sum(n []int) (r int) {
	r = 0; for _,i := range n { r += i }; return
}

func Prepend(s []int, n ...int) []int {
	return append(n, s...)
}

func Append(s []int, n ...int) []int {
	return append(s, n...)
}

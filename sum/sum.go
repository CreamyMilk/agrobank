package sum

//Ints returns the sum of integars
func Ints(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return Ints(vs[1:]...) + vs[0]
}

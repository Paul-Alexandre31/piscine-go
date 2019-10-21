package student

func IterativePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		res = nb * res
	}
	return res
}
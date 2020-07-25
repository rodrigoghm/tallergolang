package util

func Division(dividendo, divisor int) (int, int) {
	var resto int
	r := dividendo / divisor
	resto = dividendo % divisor
	return r, resto
}

func Divi(dividendo, divisor int) (r int, resto int) {
	r = dividendo / divisor
	resto = dividendo % divisor
	return
}

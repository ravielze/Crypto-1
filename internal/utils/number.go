package utils

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Modulus(d, m int) int {
	res := d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

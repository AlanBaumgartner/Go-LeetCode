package reverse

func reverse(x int) int {
	ret := 0
	if x > 0 {
		for x > 0 {
			remainder := x % 10
			ret *= 10
			ret += remainder
			x /= 10
		}
	} else {
		for x < 0 {
			remainder := x % 10
			ret *= 10
			ret += remainder
			x /= 10
		}
	}
	if ret >= 2147483647 || ret <= -2147483648 {
		return 0
	}
	return ret
}

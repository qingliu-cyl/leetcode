package main

// 超时了
func myPow(x float64, n int) float64 {
	if n == 1 {
		return x
	}

	if n == -1 {
		return 1 / x
	}

	if n == 0 {
		return 1
	}

	v := myPow(x, n/2)
	if n%2 == 0 {
		return v * v
	} else {
		if n > 0 {
			return v * v * x
		}
		return v * v * 1/x
	}
}

// 官解
//func myPow(x float64, n int) float64 {
//	if n >= 0 {
//		return quickMul(x, n)
//	}
//	return 1.0 / quickMul(x, -n)
//}
//
//func quickMul(x float64, n int) float64 {
//	if n == 0 {
//		return 1
//	}
//	y := quickMul(x, n/2)
//	if n%2 == 0 {
//		return y * y
//	}
//	return y * y * x
//}

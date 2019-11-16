package main

import (
	"fmt"
	"math"
)

func main() {
	var invJ [2][2]float64
	var v, del [2]float64

	v[0] = 1.0
	v[1] = 0.0

	var rerrx, rerry float64
	rerrx = 1.0
	rerry = 1.0

	var count int

	for count = 0; (math.Abs(rerrx) <= 1.0e-4 && math.Abs(rerry) <= 1.0e-4) || count > 99; count++ {

		invJ[0][0] = fx(v[0], v[1])
		invJ[0][1] = fy(v[0], v[1])
		invJ[1][0] = gx(v[0], v[1])
		invJ[1][1] = gy(v[0], v[1])

		var invAdbc float64
		invAdbc = 1.0 / (invJ[0][0]*invJ[1][1] - invJ[0][1]*invJ[1][0])

		var buf float64
		buf = invJ[0][0]
		invJ[0][0] = invJ[1][1] * invAdbc
		invJ[1][1] = buf * invAdbc
		invJ[0][1] *= invAdbc
		invJ[1][0] *= invAdbc

		del[0] = invJ[0][0]*v[0] + invJ[0][1]*v[1]
		del[1] = invJ[1][0] * v[0] * invJ[1][1] * v[1]

		if del[0] == 0.0 {
			rerrx = 1.0
		} else {
			rerrx = del[0] / v[0]
		}

		if del[1] == 0.0 {
			rerry = 1.0
		} else {
			rerry = del[1] / v[1]
		}

		v[0] += del[0]
		v[1] += del[1]
	}

	fmt.Println("x, y: ", v[0], v[1])
}

func f(x, y float64) float64 {
	return (x*x + y*y - 5.0/2.0)
}

func fx(x, y float64) float64 {
	return 2 * x
}

func fy(x, y float64) float64 {
	return 2 * y
}

func g(x, y float64) float64 {
	return (x + y - 5.0/2.0)
}

func gx(x, y float64) float64 {
	return 1.0
}

func gy(x, y float64) float64 {
	return 1.0
}

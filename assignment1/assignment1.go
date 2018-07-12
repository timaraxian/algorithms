package algorithms

import "fmt"

func IntMult(input1 int, input2 int) int {
	if input1 < 10 && input2 < 10 {
		return input1 * input2
	}

	a, b, c, d, n := Splitter2(input1, input2)

	p1 := IntMult(a, c) << n
	p2 := (IntMult(a, d) + IntMult(b, c)) << (n >> 1)
	p3 := IntMult(b, d)

	return p1 + p2 + p3
}

func Splitter2(input1, input2 int) (a, b, c, d int, n uint) {
	numDigits1 := numDig(input1)
	numDigits2 := numDig(input2)

	if numDigits1 > numDigits2 {
		n = numDigits1
	} else {
		n = numDigits2
	}

	if n&1 == 1 { // i.e. odd
		n++
	}

	mask := 1<<(n/2) - 1
	a = input1 >> (n / 2)
	b = mask & input1
	c = input2 >> (n / 2)
	d = mask & input2

	return a, b, c, d, n
}

func numDig(n int) (times uint) {
	times = 0
	for n != 0 {
		n = n >> 1
		times++
	}
	return times
}

func Nativemult(i1, i2 int) int {
	fmt.Print()
	return i1 * i2
}

func IterativeDoubleHalf(a int, n int) int {
	acc := 0

	for {
		if n&1 == 1 {
			acc += a

			if n == 1 {
				return acc
			}
		}

		a = a << 1
		n = n >> 1
	}
}

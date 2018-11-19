/*
Package sequence implements various mathematical sequences.
Currently implemented only Tribonacci sequence.
*/
package sequence

import (
	"errors"
	"math/big"
)

// ErrArgNotPositive returns when n argument is not positive
var ErrArgNotPositive = errors.New("Argument is not positive")

// Tribonacci returns n-th tribonacci sequence number (eg, 44 for n = 10)
func Tribonacci(n *big.Int) (*big.Int, error) {
	if n.Cmp(big.NewInt(75)) == 1 {
		return tribonacciLinearBig(n)
	}

	res, err := tribonacciLinear(int(n.Int64()))
	if err != nil {
		return nil, err
	}
	return big.NewInt(res), nil
}

func tribonacciLinear(n int) (int64, error) {
	if n <= 0 {
		return 0, ErrArgNotPositive
	}

	if n == 1 || n == 2 {
		return 0, nil
	}

	if n == 3 {
		return 1, nil
	}

	var first, second, third, res int64 = 0, 0, 1, 0
	for i := 4; i <= n; i++ {
		res = first + second + third
		first = second
		second = third
		third = res
	}

	return res, nil
}

func tribonacciLinearBig(n *big.Int) (*big.Int, error) {

	if n.Cmp(big.NewInt(0)) != 1 {
		return nil, ErrArgNotPositive
	}

	res := new(big.Int)
	if n.Cmp(big.NewInt(1)) == 0 || n.Cmp(big.NewInt(2)) == 0 {
		return res, nil
	}
	if n.Cmp(big.NewInt(3)) == 0 {
		res.SetInt64(1)
		return res, nil
	}

	one := big.NewInt(1)
	var first, second, third = new(big.Int), new(big.Int), big.NewInt(1)
	for i := new(big.Int).SetInt64(3); i.Cmp(n) == -1; i.Add(i, one) {
		res.Add(first, second)
		res.Add(res, third)
		first.Set(second)
		second.Set(third)
		third.Set(res)
	}
	return res, nil
}

package sequence

import (
	"math/big"
	"testing"
)

var tests = []struct {
	n        int
	expected int64
}{
	{1, 0},
	{2, 0},
	{3, 1},
	{4, 1},
	{20, 19513},
	{38, 1132436852},
	{75, 7015254043203144209},
}

func TestTribonacci(t *testing.T) {
	for _, tt := range tests {
		current, err := Tribonacci(big.NewInt(int64(tt.n)))
		if err != nil {
			if tt.n > 0 {
				t.Error(err)
			}
		}

		if big.NewInt(tt.expected).Cmp(current) != 0 {
			t.Errorf("Expected: %d, got: %d", tt.expected, current)
		}
	}
}

func TestTribonacci_fail(t *testing.T) {
	_, err := Tribonacci(big.NewInt(0))

	if err != ErrArgNotPositive {
		t.Errorf("Wrong result: %v", err)
	}

	_, err = Tribonacci(big.NewInt(-1))

	if err != ErrArgNotPositive {
		t.Errorf("Wrong result: %v", err)
	}
}

func TestLinear(t *testing.T) {
	for _, tt := range tests {
		current, err := tribonacciLinear(tt.n)
		if err != nil {
			if tt.n > 0 {
				t.Error(err)
			}
		}

		if current != tt.expected {
			t.Errorf("Expected: %d, got: %d", tt.expected, current)
		}
	}
}

func TestLinearBig(t *testing.T) {
	for _, tt := range tests {
		current, err := tribonacciLinearBig(big.NewInt(int64(tt.n)))
		if err != nil {
			if tt.n > 0 {
				t.Error(err)
			}
		}

		if big.NewInt(tt.expected).Cmp(current) != 0 {
			t.Errorf("Expected: %d, got: %d", tt.expected, current)
		}
	}
}

func BenchmarkTribonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(big.NewInt(10))
	}
}

func BenchmarkTribonacci_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(big.NewInt(100))
	}
}

func BenchmarkTribonacci_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(big.NewInt(1000))
	}
}

func BenchmarkTribonacci_1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(big.NewInt(1000000))
	}
}

func BenchmarkLinear_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacciLinear(10)
	}
}

func BenchmarkLinearBig_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacciLinearBig(big.NewInt(10))
	}
}

func BenchmarkLinearBig_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacciLinearBig(big.NewInt(100))
	}
}

func BenchmarkLinearBig_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacciLinearBig(big.NewInt(1000))
	}
}

func BenchmarkLinearBig_1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tribonacciLinearBig(big.NewInt(1000000))
	}
}

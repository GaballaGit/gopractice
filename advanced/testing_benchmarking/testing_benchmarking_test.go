package advanced

import (
	"math/rand"
	"testing"
)

// Note: go test -bench=. -memprofile mem.pprof file.go
// You can observce this with go tool pprof file.pprof
func GenerateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}

	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	size := 100
	slice := GenerateRandomSlice(size)
	if len(slice) != size {
		t.Errorf("expected slice size %d, recieved %d", size, len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for range b.N {
		GenerateRandomSlice(1000)
	}
}

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(1000)
	b.ResetTimer()
	for range b.N {
		SumSlice(slice)
	}
}

/*func Add(a int, b int) int {
	return a + b
}

// ============================= Benchmarking =============================
func BenchmarkAddSmallInput(b *testing.B) {
	for range b.N {
		Add(2, 3)
	}
}

func BenchmarkAddMediumInput(b *testing.B) {
	for range b.N {
		Add(200, 300)
	}
}

func BenchmarkAddBigInput(b *testing.B) {
	for range b.N {
		Add(200000000, 300000000)
	}
}*/

// ============================= Testing =============================

/*func TestAddSubtests(t *testing.T) {
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 1, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
			result := Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("result = %d; want %d", result, test.expected)
			}
		})
	}
}*/

/*func TestAddTableDriven(t *testing.T) {
	// Test cases stored in the struct... pretty smart.
	tests := []struct{ a, b, expected int }{
		{2, 3, 5},
		{0, 1, 0},
		{-1, 1, 0},
	}

	for _, test := range tests {
		result := Add(test.a, test.b)
		if result != test.expected {
			t.Errorf("Add(%d, %d) = %d: want %d\n", test.a, test.b, result, test.expected)
		}
	}
}*/

/*func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2,3) = %d want %d\n", result, expected)
	}
}*/

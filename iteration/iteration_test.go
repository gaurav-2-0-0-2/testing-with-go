package iteration

import ( 
	"fmt"
	"testing"
) 

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 3)
	fmt.Println(result)
	// Output: aaa
}

// Benchmarking
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("given a negative repeat count, default to 5", func(t *testing.T) {
		repeated := Repeat("a", -10)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but received %q", expected, repeated)
		}
	})

	t.Run("given a repeat count of 22", func(t *testing.T) {
		repeated := Repeat("a", 22)
		expected := "aaaaaaaaaaaaaaaaaaaaaa"

		if repeated != expected {
			t.Errorf("expected %q but received %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 2)
	fmt.Println(repeated)
	// Output: aa
}

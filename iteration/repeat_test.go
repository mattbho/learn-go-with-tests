package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeats the String N times", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("Repeats the String N times using built in String.Repeat", func(t *testing.T) {
		repeated := StringRepeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeat := Repeat("h", 5)
	fmt.Println(repeat)
	// Output: hhhhh
}

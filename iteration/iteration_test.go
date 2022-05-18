package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("repeats 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrect(t, got, want)
	})

	t.Run("repeats 12 times", func(t *testing.T) {
		got := Repeat("yo", 12)
		want := "yoyoyoyoyoyoyoyoyoyoyoyo"
		assertCorrect(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("hi", 10)
	fmt.Println(repeated)
	// Output: hihihihihihihihihihi
}

package iteration

import (
	"fmt"
	"testing"
)

var repeatTests = []struct {
	in    string
	count int
	out   string
}{
	{"a", 5, "aaaaa"},
	{"b", 4, "bbbb"},
	{"é", 9, "ééééééééé"},
}

func TestFlagParser(t *testing.T) {
	for _, tt := range repeatTests {
		testMessage := fmt.Sprintf("Testing char '%s', with count %d", tt.in, tt.count)
		t.Run(testMessage, func(t *testing.T) {
			r := Repeat(tt.in, tt.count)
			if r != tt.out {
				t.Errorf("got %q, want %q", r, tt.out)
			}
		})
	}
}

// func TestRepeat(t *testing.T) {
// 	repeated := Repeat("a", 5)
// 	expected := "aaaaa"
//
// 	if repeated != expected {
// 		t.Errorf("expected '%s' but got '%s'", expected, repeated)
// 	}
// }
//
// func ExampleRepeat() {
// 	fmt.Println(Repeat("b", 6))
// 	// Output: bbbbbb
// }
//
// func BenchmarkRepeat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Repeat("a", 5)
// 	}
// }

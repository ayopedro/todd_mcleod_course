package sayings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Ayotunde")

	if s != "Welcome my dear Ayotunde" {
		t.Error("got", s, "expected", "Welcome my dear Ayotunde")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output
	// Welcome my dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Ayotunde")
	}
}

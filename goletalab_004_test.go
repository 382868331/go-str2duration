package str2duration

import (
	"testing"
	"time"
)

func TestGoletaStr2Duration004LargeNanoseconds(t *testing.T) {
	const input = "1000000000000000000ns"
	got, err := ParseDuration(input)
	if err != nil {
		t.Fatalf("ParseDuration(%q) returned %v", input, err)
	}
	want := time.Duration(1_000_000_000_000_000_000)
	if got != want {
		t.Fatalf("ParseDuration(%q)=%v,want=%v", input, got, want)
	}
}

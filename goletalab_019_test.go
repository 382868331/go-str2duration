package str2duration
import("testing";"time")
func TestGoletaStr2Duration019(t *testing.T) {
	got, err := ParseDuration("1s9ms")
	want, _ := time.ParseDuration("1009ms")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration019AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("2m9s")
	want, _ := time.ParseDuration("129s")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

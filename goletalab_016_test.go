package str2duration
import("testing";"time")
func TestGoletaStr2Duration016(t *testing.T) {
	got, err := ParseDuration("9s")
	want, _ := time.ParseDuration("9s")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration016AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("19ms")
	want, _ := time.ParseDuration("19ms")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

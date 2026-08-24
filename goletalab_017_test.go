package str2duration
import("testing";"time")
func TestGoletaStr2Duration017(t *testing.T) {
	got, err := ParseDuration("1.5s")
	want, _ := time.ParseDuration("1500ms")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration017AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("0.25m")
	want, _ := time.ParseDuration("15s")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

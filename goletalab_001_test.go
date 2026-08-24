package str2duration
import("testing";"time")
func TestGoletaStr2Duration001(t *testing.T) {
	got, err := ParseDuration("1d")
	want, _ := time.ParseDuration("24h")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration001AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("2d")
	want, _ := time.ParseDuration("48h")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

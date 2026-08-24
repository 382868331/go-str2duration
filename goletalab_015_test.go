package str2duration
import("testing";"time")
func TestGoletaStr2Duration015(t *testing.T) {
	got, err := ParseDuration("0")
	want, _ := time.ParseDuration("0s")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration015AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("+0")
	want, _ := time.ParseDuration("0s")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

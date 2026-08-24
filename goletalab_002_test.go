package str2duration
import("testing";"time")
func TestGoletaStr2Duration002(t *testing.T) {
	got, err := ParseDuration("1w")
	want, _ := time.ParseDuration("168h")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

func TestGoletaStr2Duration002AdjacentBoundary(t *testing.T) {
	got, err := ParseDuration("2w")
	want, _ := time.ParseDuration("336h")
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	if got != want { t.Fatalf("got=%v want=%v", got, want) }
}

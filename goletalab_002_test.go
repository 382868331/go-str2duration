package str2duration
import("testing";"time")
func TestGoletaStr2Duration002(t *testing.T) {
	got, err := ParseDuration("1w")
	want, _ := time.ParseDuration("168h")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

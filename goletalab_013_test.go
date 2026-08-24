package str2duration
import("testing";"time")
func TestGoletaStr2Duration013(t *testing.T) {
	got, err := ParseDuration("-2s")
	want, _ := time.ParseDuration("-2s")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

package str2duration
import("testing";"time")
func TestGoletaStr2Duration019(t *testing.T) {
	got, err := ParseDuration("1s9ms")
	want, _ := time.ParseDuration("1009ms")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

package str2duration
import("testing";"time")
func TestGoletaStr2Duration003(t *testing.T) {
	got, err := ParseDuration("2ns")
	want, _ := time.ParseDuration("2ns")
	if err != nil || got != want { t.Fatalf("got=%v want=%v err=%v", got, want, err) }
}

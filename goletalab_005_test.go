package str2duration

import "testing"

func TestGoletaStr2Duration005QuotedError(t *testing.T) {
    _, err := ParseDuration("1x")
    const want = `time: unknown unit "x" in duration "1x"`
    if err == nil || err.Error() != want { t.Fatalf("error=%v,want=%s", err, want) }
}

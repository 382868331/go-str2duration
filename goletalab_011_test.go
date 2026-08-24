package str2duration

import "testing"

func TestGoletaStr2Duration011EmptyInput(t *testing.T) {
    if got, err := ParseDuration(""); err == nil { t.Fatalf("got=%v,want invalid-duration error", got) }
}

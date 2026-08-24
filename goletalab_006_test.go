package str2duration

import "testing"

func TestGoletaStr2Duration006ZeroString(t *testing.T) {
    if got := String(0); got != "0s" { t.Fatalf("String(0)=%q,want=0s", got) }
}

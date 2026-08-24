package str2duration

import ("testing"; "time")

func TestGoletaStr2Duration007NegativeString(t *testing.T) {
    if got := String(-2*time.Second); got != "-2s" { t.Fatalf("String(-2s)=%q,want=-2s", got) }
}

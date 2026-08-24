package str2duration

import ("testing"; "time")

func TestGoletaStr2Duration009DecimalDigits(t *testing.T) {
    if got := String(12*time.Second); got != "12s" { t.Fatalf("String(12s)=%q,want=12s", got) }
}

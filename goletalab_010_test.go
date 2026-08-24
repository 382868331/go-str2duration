package str2duration

import ("testing"; "time")

func TestGoletaStr2Duration010LongFraction(t *testing.T) {
    got, err := ParseDuration("0.123456789012345678901s")
    if err != nil || got != 123456789*time.Nanosecond { t.Fatalf("got=%v,err=%v", got, err) }
}

package str2duration

import ("testing"; "time")

func TestGoletaStr2Duration012ZeroComponent(t *testing.T) {
    got, err := ParseDuration("1h0m")
    if err != nil || got != time.Hour { t.Fatalf("got=%v,err=%v,want=1h", got, err) }
}

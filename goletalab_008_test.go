package str2duration

import ("testing"; "time")

func TestGoletaStr2Duration008MaximumString(t *testing.T) {
    const want = "15250w1d23h47m16s854ms775us807ns"
    if got := String(time.Duration(1<<63-1)); got != want { t.Fatalf("String(MaxInt64)=%q,want=%q", got, want) }
}

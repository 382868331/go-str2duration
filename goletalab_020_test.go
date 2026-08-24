package str2duration

import "testing"

func TestGoletaStr2Duration020CompoundOverflow(t *testing.T){if got,err:=ParseDuration("9223372036854775807ns1ns");err==nil{t.Fatalf("got=%v,want overflow error",got)}}

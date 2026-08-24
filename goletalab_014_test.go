package str2duration

import "testing"

func TestGoletaStr2Duration014UppercaseUnit(t *testing.T){if got,err:=ParseDuration("1S");err==nil{t.Fatalf("got=%v,want unknown-unit error",got)}}

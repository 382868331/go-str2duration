package str2duration

import "testing"

func TestGoletaStr2Duration018FractionTruncation(t *testing.T){got,err:=ParseDuration("0.9ns");if err!=nil||got!=0{t.Fatalf("got=%v,err=%v,want=0",got,err)}}

package models

import "testing"

func TestQuery(t *testing.T) {
	s:=Query("1","雷军")
	if s==nil||len(s)<1 {
		t.Errorf("错误，长度%d",len(s))
	}
}
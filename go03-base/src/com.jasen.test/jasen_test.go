package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	type aryList struct{
		a int
		b int
	}

	testValues := []aryList{
		{1,2},{454,11},{55,200},
	}

	for _,value := range testValues{
		sum := Add(value.a,value.b)
		if sum <= 300 {
			t.Logf("[%d,%d]the result lt 300",value.a,value.b)
		} else {
			t.Fatalf("[%d,%d]the result gt 300",value.a,value.b)
		}
	}
}

func TestAdd1(t *testing.T) {
	t.Error("the result is error")
}

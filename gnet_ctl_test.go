package main

import (
	"reflect"
	"testing"
)

func TestAsStringTo32(t *testing.T) {
	assertEqual := func(v interface{}, v1 interface{}) {
		if v != v1 {
			t.Errorf("the values %v and %v are not equal.", v1, v)
		}
	}
	var n uint32
	var str string = "650001"
	ustr := StrToUin32(str)
	us := reflect.ValueOf(ustr).Kind()
	u := reflect.ValueOf(n).Kind()
	assertEqual(us, u)
}

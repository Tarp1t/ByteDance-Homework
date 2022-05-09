package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	//定义一个测试用例类型
	type test struct {
		input string
		sep   string
		want  []string
	}
	//定义一个存储测试用例的切片
	tests := []test{
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "这是一次测试", sep: "一", want: []string{"这是", "次测试"}},
	}
	for _, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Excepted:%v,got:%v", tc.want, got)
		}
	}
}

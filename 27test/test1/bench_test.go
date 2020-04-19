package main

import "testing"

//压力测试
//BenchmarkAdd-4   	2000000000 代表执行的次数
// 0.44 ns/op  每次平均执行时间为 0.44纳秒

func BenchmarkAdd(b *testing.B){
	//N ：是由自动化框架自己确定的
	for i:=0;i<b.N;i++  {
		a:=1
		b:=2
		Add(a,b)
	}
}

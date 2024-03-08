package main

import (
	"testing"
)

type BigStruct struct {
	Data [1 << 24]int // このサイズは約16MB
}

// 値として渡す関数
func processValue(s BigStruct) int {
	return s.Data[0]
}

// ポインタとして渡す関数
func processPointer(s *BigStruct) int {
	return s.Data[0]
}

// ベンチマークテスト関数
func BenchmarkValue(b *testing.B) {
	s := BigStruct{}
	for i := 0; i < b.N; i++ {
		processValue(s)
	}
}

// ポインタとして渡すバージョンのベンチマーク
func BenchmarkPointer(b *testing.B) {
	s := BigStruct{}
	for i := 0; i < b.N; i++ {
		processPointer(&s)
	}
}

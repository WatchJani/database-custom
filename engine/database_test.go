package engine

import (
	"root/index"
	"testing"
)

// 6.066 ns/op
func BenchmarkCreateDataBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewDataBase("yovan")
	}
}

// 62.43 ns/op
func BenchmarkCreateTable(b *testing.B) {
	system := NewDataBase("system")

	for i := 0; i < b.N; i++ {
		system.CreateTable("user")
	}
}

// 63.74 ns/op
func BenchmarkAddIndex(b *testing.B) {
	system := NewDataBase("yovan")
	system.CreateTable("user")

	for i := 0; i < b.N; i++ {
		system.AddIndex("user", index.HASH_INDEX)
	}
}

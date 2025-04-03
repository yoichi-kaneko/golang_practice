// ベンチマークテストの検証
package main

import (
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := decode("post.json")
		if err != nil {
			b.Errorf("Error decoding JSON: %v", err)
			return
		}
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := unmarshal("post.json")
		if err != nil {
			b.Errorf("Error unmarshaling JSON: %v", err)
			return
		}
	}
}

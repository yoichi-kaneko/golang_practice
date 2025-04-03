// main.goのテスト
package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Errorf("Error decoding JSON: %v", err)
		return
	}

	if post.Id != 1 {
		t.Errorf("Expected Id 1, got %d", post.Id)
	}
	if post.Content != "Hello, world!" {
		t.Errorf("Expected Content 'Hello, world!', got '%s'", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encode test for now")
}

func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second) // Simulate a long-running test
}

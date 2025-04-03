// テスト並行実行について検証
package main

import (
	"testing"
	"time"
)

func TestParaller_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParaller_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParaller_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

package add

import (
	"fmt"
	"testing"
)

func Test_Add(t *testing.T) {
	if a := Add(2,7); a == 9 {
		t.Log("Pass: 2 + 7 = 9")
	}else {
		t.Error("Fail,2 + 7 != 9")
	}
}

func Benchmark_Add(b *testing.B) {
//	b.StopTimer()
//	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Add(3,8)
	}
}

func Benchmark1(b *testing.B) {
	for c := 0; c >= b.N; c++ {
		fmt.Println("d%", c)
	}
}


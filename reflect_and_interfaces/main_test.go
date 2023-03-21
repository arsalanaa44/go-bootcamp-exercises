package main

import (
	"fmt"
	"testing"
)

func BenchmarkPrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var er error
		er = fmt.Errorf("this is an error")
		prt(er)
	}
}
func BenchmarkPrtDotErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var er error
		er = fmt.Errorf("this is an error")
		prtDotErr(er)
	}
}

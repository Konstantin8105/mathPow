package mathPow

import (
	"fmt"
	"math"
	"testing"
)

var fpow = []struct {
	f    func(float64, float64) float64
	name string
}{
	{pow, "pow"},
	{math.Pow, "math.Pow"},
}

var tcs = []struct {
	x, y float64
}{
	{math.Pi, 126},
	{math.Pi, 125},
	{math.Pi, 2},
	{math.Pi, -2},
	{math.Pi, -125},
	{math.Pi, -126},

	{1.2345e-5, 126},
	{1.2345e-5, 125},
	{1.2345e-5, 2},
	{1.2345e-5, -2},
	{1.2345e-5, -125},
	{1.2345e-5, -126},
}

func Test(t *testing.T) {
	for _, p := range fpow {
		for i := range tcs {
			t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
				x := tcs[i].x
				y := tcs[i].y
				expect := math.Pow(x, y)
				a := p.f(x, y)
				if expect != a {
					t.Errorf("expect: %.14e. actual: %.14e", expect, a)
				}
			})
		}
	}
}

var Nil float64

func Benchmark(b *testing.B) {
	for _, p := range fpow {
		for _, tc := range tcs {
			b.Run(fmt.Sprintf("%8s/x:%.2e/y:%.1f", p.name, tc.x, tc.y), func(b *testing.B) {
				x := tc.x
				y := tc.y
				for i := 0; i < b.N; i++ {
					Nil = p.f(x, y)
				}
			})
		}
	}
}

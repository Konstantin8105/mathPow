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
	{pow_des1, "pow_des1"},
	{pow_des2, "pow_des2"},
	{pow_sq1, "pow_sq1"},
	{pow_win, "pow_win"},
	{math.Pow, "math.Pow"},
}

var tcs = []struct {
	name string
	x, y float64
}{
	{"SIMPLE", math.Pi, 7.0},
	{"SQRT2+", math.Pi, 1.0 / 2.0},
	{"SQRT2-", math.Pi, -1.0 / 2.0},
	{"SQRT3", math.Pi, -1.0 / 3.0},
	{"SQRT8+", math.Pi, 1.0 / 8.0},
	{"SQRT8-", math.Pi, -1.0 / 8.0},
	{"SQRT32+", math.Pi, 1.0 / 32.0},
	{"SQRT32-", math.Pi, -1.0 / 32.0},
	{"SQRT64+", math.Pi, 1.0 / 64.0},
	{"SQRT64-", math.Pi, -1.0 / 64.0},
	{"BASE", 1.2345e-5, -2.5},

	{"PI", math.Pi, 126},
	{"PI", math.Pi, 125},
	{"PI", math.Pi, 2},
	{"PI", math.Pi, -2},
	{"PI", math.Pi, -125},
	{"PI", math.Pi, -126},

	{"SMALL", 1.2345e-5, 126},
	{"SMALL", 1.2345e-5, 125},
	{"SMALL", 1.2345e-5, 2},
	{"SMALL", 1.2345e-5, -2},
	{"SMALL", 1.2345e-5, -125},
	{"SMALL", 1.2345e-5, -126},
}

func Test(t *testing.T) {
	tcs := tcs
	for i := -200; i <= 200; i++ {
		tcs = append(tcs, struct {
			name string
			x, y float64
		}{"e", math.Pi, float64(i)})
	}
	for i := -200; i <= 200; i++ {
		tcs = append(tcs, struct {
			name string
			x, y float64
		}{"e", math.Pi, float64(i) + 0.1})
	}
	for _, p := range fpow {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%8s/x:%.2e/y:%.1f", tc.name, p.name, tc.x, tc.y), func(t *testing.T) {
				x := tc.x
				y := tc.y
				expect := math.Pow(x, y)
				a := p.f(x, y)
				if math.Abs((expect-a)/a) > 1e-15 {
					t.Errorf("\nexpect: %.14e\nactual: %.14e\ndiff  : %.14e", expect, a, (expect-a)/a)
				}
			})
		}
	}
}

var Nil float64

func Benchmark(b *testing.B) {
	for _, p := range fpow {
		for _, tc := range tcs {
			b.Run(fmt.Sprintf("%s/%8s/x:%.2e/y:%.1f", tc.name, p.name, tc.x, tc.y), func(b *testing.B) {
				x := tc.x
				y := tc.y
				for i := 0; i < b.N; i++ {
					Nil = p.f(x, y)
				}
			})
		}
	}
}

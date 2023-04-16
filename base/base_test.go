package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a T
		b T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			args: args[int]{
				a: 0,
				b: 1,
			},
			want: 0,
		},
		{
			args: args[int]{
				a: 1,
				b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Min(tt.args.a, tt.args.b), "Min(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a T
		b T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			args: args[int]{
				a: 1,
				b: 0,
			},
			want: 1,
		},
		{
			args: args[int]{
				a: 0,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Max(tt.args.a, tt.args.b), "Max(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestCmpOrdered(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a T
		b T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			args: args[int]{
				a: 0,
				b: 1,
			},
			want: -1,
		},
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
				b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CmpOrdered(tt.args.a, tt.args.b), "CmpOrdered(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestCmpComparable(t *testing.T) {
	type args[T comparable] struct {
		a T
		b T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{
			args: args[int]{
				a: 0,
				b: 1,
			},
			want: 1,
		},
		{
			args: args[int]{
				a: 0,
				b: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CmpComparable(tt.args.a, tt.args.b), "CmpComparable(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func TestSliceCmp(t *testing.T) {
	cmp := SliceCmp(CmpOrdered[int])
	assert.Equal(t, -1, cmp([]int{0}, []int{1}))
	assert.Equal(t, -1, cmp(nil, []int{0}))
}

func TestReverseCmp(t *testing.T) {
	cmp := ReverseCmp(CmpOrdered[int])
	assert.Equal(t, 1, cmp(0, 1))
}

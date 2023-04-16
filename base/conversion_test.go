package base

import "testing"

func TestBool2Int(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				b: true,
			},
			want: 1,
		},
		{
			args: args{
				b: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool2Int(tt.args.b); got != tt.want {
				t.Errorf("Bool2Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

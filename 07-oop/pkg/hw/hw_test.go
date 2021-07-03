package hw

import "testing"

func TestDistance(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "#1",
			args: args{1,1,4,5},
			want: 5.0,
		},
		{
			name: "#2",
			args: args{0,0,0,0},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}
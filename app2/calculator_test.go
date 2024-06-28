package main

import "testing"

func TestAddition(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "Addition of 2 and 3",
			a:    2,
			b:    3,
			want: 7,
		},
		{
			name: "Addition of 5 and 3",
			a:    5,
			b:    3,
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.a, tt.b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Addition(2, 3)
	}
}

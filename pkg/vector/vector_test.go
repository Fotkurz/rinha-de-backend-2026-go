package vector_test

import (
	"fmt"
	"testing"

	"github.com/Fotkurz/rinha-de-backend-2026-go/pkg/vector"
)

func TestNormalize(t *testing.T) {
	type args struct {
		v, coef float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"1.0 if value is over 1.0", args{3, 1}, 1.0},
		{"0.0 if value is below 0.0", args{-1, 1}, 0.0},
		{"value if result between 0 and 1", args{1, 2}, 0.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vector.Normalize(tt.args.v, tt.args.coef)
			if tt.want != got {
				t.Errorf("expected=(%v), but got=(%v)", tt.want, got)
			}
		})
	}
}

func TestEuclidian(t *testing.T) {
	type args struct {
		v1, v2 []float32
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{"success", args{v1: []float32{1.0, 0.96, 0.96}, v2: []float32{0.5796, 0.9167, 1.00}}, 0.424513},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := vector.EuclidianDistance(tt.args.v1, tt.args.v2)
			if fmt.Sprintf("%.4f", tt.want) != fmt.Sprintf("%.4f", got) {
				t.Errorf("Expected=(%f) but got=(%f)", tt.want, got)
			}
		})
	}
}

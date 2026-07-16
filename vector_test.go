package mathrock

import (
	"reflect"
	"testing"
)

func TestVectorDot(t *testing.T) {
	tests := []struct {
		name    string
		a       Vector
		b       Vector
		want    float64
		wantErr bool
	}{
		{
			name:    "equal length vectors",
			a:       Vector{1, 2, 3},
			b:       Vector{4, 5, 6},
			want:    32,
			wantErr: false,
		},
		{
			name:    "mismatched length vectors",
			a:       Vector{1, 2},
			b:       Vector{1, 2, 3},
			want:    0,
			wantErr: true,
		},
		{
			name:    "empty vectors",
			a:       Vector{},
			b:       Vector{},
			want:    0,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Dot(tt.b)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Dot() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("Dot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorAdd(t *testing.T) {
	tests := []struct {
		name    string
		a       Vector
		b       Vector
		want    Vector
		wantErr bool
	}{
		{
			name:    "equal length vectors",
			a:       Vector{1, 2, 3},
			b:       Vector{4, 5, 6},
			want:    Vector{5, 7, 9},
			wantErr: false,
		},
		{
			name:    "mismatched length vectors",
			a:       Vector{1, 2},
			b:       Vector{1, 2, 3},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty vectors",
			a:       Vector{},
			b:       Vector{},
			want:    Vector{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Add(tt.b)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorSubtract(t *testing.T) {
	tests := []struct {
		name    string
		a       Vector
		b       Vector
		want    Vector
		wantErr bool
	}{
		{
			name:    "equal length vectors",
			a:       Vector{4, 5, 6},
			b:       Vector{1, 2, 3},
			want:    Vector{3, 3, 3},
			wantErr: false,
		},
		{
			name:    "mismatched length vectors",
			a:       Vector{1, 2},
			b:       Vector{1, 2, 3},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty vectors",
			a:       Vector{},
			b:       Vector{},
			want:    Vector{},
			wantErr: false,
		},
		{
			name:    "result has negative values",
			a:       Vector{1, 2, 3},
			b:       Vector{4, 5, 6},
			want:    Vector{-3, -3, -3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.a.Subtract(tt.b)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Subtract() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVectorScale(t *testing.T) {
	tests := []struct {
		name   string
		v      Vector
		scalar float64
		want   Vector
	}{
		{
			name:   "positive scalar",
			v:      Vector{1, 2, 3},
			scalar: 2,
			want:   Vector{2, 4, 6},
		},
		{
			name:   "negative scalar",
			v:      Vector{1, 2, 3},
			scalar: -1,
			want:   Vector{-1, -2, -3},
		},
		{
			name:   "zero scalar",
			v:      Vector{1, 2, 3},
			scalar: 0,
			want:   Vector{0, 0, 0},
		},
		{
			name:   "empty vector",
			v:      Vector{},
			scalar: 5,
			want:   Vector{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.v.Scale(tt.scalar)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scale() = %v, want %v", got, tt.want)
			}
		})
	}
}
package main

import (
	"testing"
)

func Test_format2Conv(t *testing.T) {
	tests := []struct {
		name   string
		format string
		in     string
		want   float64
	}{
		{"duration", "d", "1h2s", 60*60 + 2},
		{"minutes", "m", "2", 60 * 2},
		{"hours", "h", "0.5", 60 * 30},
		{"seconds", "s", "1", 1},
		{"milliseconds", "i", "100", 0.1},
		{"microseconds", "u", "1.1", 0.0000011},
		{"nanoseconds", "n", "890", 0.00000089},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := format2Conv(tt.format)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			got, err := f(tt.in)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if got != tt.want {
				t.Errorf("got %f, want %f", got, tt.want)
			}
		})
	}
}

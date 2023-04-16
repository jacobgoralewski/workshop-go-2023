package main

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "2+2", args: args{a: 2, b: 2}, want: 4},
		{name: "2+3", args: args{a: 2, b: 3}, want: 5},
		{name: "3+3", args: args{a: 3, b: 3}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

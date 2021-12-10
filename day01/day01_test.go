package main

import "testing"

func Test_calcFloor(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{[]byte("(((")},
			want: 3,
		},
		{
			args: args{[]byte("()()")},
			want: 0,
		},
		{
			args: args{[]byte(")())())")},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.input), func(t *testing.T) {
			if got := calcFloor(tt.args.input); got != tt.want {
				t.Errorf("calcFloor('%s') = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}

func Test_calcBasement(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: ")",
			args: args{[]byte(")")},
			want: 1,
		},
		{
			name: "()())",
			args: args{[]byte("()())")},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcBasement(tt.args.input); got != tt.want {
				t.Errorf("calcBasement('%s') = %v, want %v", tt.args.input, got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_calculateHomesVisited(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		args      args
		wantHomes int
	}{
		{
			args:      args{[]byte("<")},
			wantHomes: 2,
		},
		{
			args:      args{[]byte("^>v<")},
			wantHomes: 4,
		}, {
			args:      args{[]byte("^v^v^v^v^v")},
			wantHomes: 2,
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.input), func(t *testing.T) {
			if gotHomes := calculateHomesVisited(tt.args.input); gotHomes != tt.wantHomes {
				t.Errorf("calculateHomesVisited(%v) = %v, want %v", string(tt.args.input), gotHomes, tt.wantHomes)
			}
		})
	}
}

func Test_split(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		args       args
		wantInput1 []byte
		wantInput2 []byte
	}{
		{
			args:       args{[]byte("^>v<")},
			wantInput1: []byte("^v"),
			wantInput2: []byte("><"),
		},

		{
			args:       args{[]byte("^v^v^v^v^v")},
			wantInput1: []byte("^^^^^"),
			wantInput2: []byte("vvvvv"),
		},
	}
	for _, tt := range tests {
		t.Run(string(tt.args.input), func(t *testing.T) {
			gotInput1, gotInput2 := split(tt.args.input)
			if !reflect.DeepEqual(gotInput1, tt.wantInput1) {
				t.Errorf("split() gotInput1 = %v, want %v", gotInput1, tt.wantInput1)
			}
			if !reflect.DeepEqual(gotInput2, tt.wantInput2) {
				t.Errorf("split() gotInput2 = %v, want %v", gotInput2, tt.wantInput2)
			}
		})
	}
}

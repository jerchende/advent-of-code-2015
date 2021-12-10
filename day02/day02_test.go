package main

import "testing"

func Test_paper(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args                  args
		wantPaper, wantRibbon int
	}{
		{
			args:       args{"2x3x4"},
			wantPaper:  58,
			wantRibbon: 34,
		},
		{
			args:       args{"1x1x10"},
			wantPaper:  43,
			wantRibbon: 14,
		},
		{
			args:       args{"1x10x1"},
			wantPaper:  43,
			wantRibbon: 14,
		},
		{
			args:       args{"4x3x2"},
			wantPaper:  58,
			wantRibbon: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.input, func(t *testing.T) {
			if gotPaper, gotRibbon := forPackage(tt.args.input); gotPaper != tt.wantPaper || gotRibbon != tt.wantRibbon {
				t.Errorf("forPackage() = %v, %v, want %v, %v", gotPaper, gotRibbon, tt.wantPaper, tt.wantRibbon)
			}
		})
	}
}

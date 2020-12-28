package main

import "testing"

func Test_phoneValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{"+55 11 98888-8888"}, false},
		{"T2", args{"11 98888-8888"}, false},
		{"T3", args{"988888888"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := phoneValidate(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("phoneValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

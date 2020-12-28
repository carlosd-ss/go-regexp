package main

import "testing"

func TestCnpjValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{"00000000000000"}, false},
		{"T2", args{"00.000.000/0000-00"}, false},
		{"T3", args{"00-000-000-0000-00"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CnpjValidate(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CnpjValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

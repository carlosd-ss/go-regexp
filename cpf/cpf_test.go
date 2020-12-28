package main

import "testing"

func TestCpfValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{"00000000000"}, false},
		{"T2", args{"000-000-000-00"}, false},
		{"T3", args{"000.000.000.00"}, false},
		{"T4", args{"000.000.000-00"}, false},
		{"T5", args{"000-000-000.00"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CpfValidate(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CpfValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package main

import "testing"

func TestCepValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{"00000000"}, false},
		{"T3", args{"00000 000"}, false},
		{"T4", args{"00000-000"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CepValidate(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("CepValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

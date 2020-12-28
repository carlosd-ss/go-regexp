package main

import "testing"

func Test_emailValidate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"T1", args{"gopher@gmail.com"}, false},
		{"T2", args{"gopher-go@outlook.com"}, false},
		{"T3", args{"gopher-br@hotmail.com.br"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := emailValidate(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("emailValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

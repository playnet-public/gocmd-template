package template

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"abc", args{"abc"}, "this is some template project called abc"},
		{"hello", args{"hello"}, "this is some template project called hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Template(tt.args.name); got != tt.want {
				t.Errorf("Template() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"abc", "ups", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Error()
			if (err != nil) != tt.wantErr {
				t.Errorf("Error() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

package writer

import (
	"testing"
)

func TestWithLabel(t *testing.T) {
	type args struct {
		raw   string
		label string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "with label",
			args: args{
				raw:   "mixin://codes/code",
				label: "Lock 1 BTC",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if raw, label := handleLabel(tt.args.raw, tt.args.label); raw != tt.args.raw || label != tt.args.label {
				t.Errorf("handleLabel() = %v,%v want %v,%v", raw, label, tt.args.raw, tt.args.label)
			}
		})
	}
}

func handleLabel(raw, label string) (string, string) {
	s := WithLabel(raw, label)
	raw, label, _ = ExtractLabel(s)
	return raw, label
}

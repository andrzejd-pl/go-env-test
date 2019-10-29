package another

import "testing"

func TestEcho(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Second", "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Echo(); got != tt.want {
				t.Errorf("Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}

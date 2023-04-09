package goos

import (
	"runtime"
	"testing"
)

func TestIsSpecifiedPlatform(t *testing.T) {
	type args struct {
		platform string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: IsDarwin, args: args{platform: IsDarwin}, want: false},
		{name: IsLinux, args: args{platform: IsLinux}, want: false},
		{name: IsWindows, args: args{platform: IsWindows}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == runtime.GOOS {
				tt.want = true
			}
			if got := IsSpecifiedPlatform(tt.args.platform); got != tt.want {
				t.Errorf("IsSpecifiedPlatform() = %v, want %v", got, tt.want)
			}
		})
	}
}

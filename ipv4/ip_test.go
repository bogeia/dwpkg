package ipv4

import (
	"testing"
)

func TestIp2long(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			args: struct{ ip string }{ip: "127.0.0.1"},
			want: 2130706433,
		},
		{
			args: struct{ ip string }{ip: "10.10.10.1"},
			want: 168430081,
		},
		{
			args: struct{ ip string }{ip: "10.10.10.256"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertIpToUint32(tt.args.ip); got != tt.want {
				t.Errorf("Ip2long() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLong2ip(t *testing.T) {
	type args struct {
		long uint32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: struct{ long uint32 }{long: 168430081},
			want: "10.10.10.1",
		},
		{
			args: struct{ long uint32 }{long: 2130706433},
			want: "127.0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertUint32ToIp(tt.args.long); got != tt.want {
				t.Errorf("Long2ip() = %v, want %v", got, tt.want)
			}
		})
	}
}

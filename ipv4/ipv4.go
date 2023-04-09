package ipv4

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"

	"github.com/bogeia/dwpkg/regexpx"
)

// ValidateIpv4 checks whether given `ip` a valid IPv4 address.
func ValidateIpv4(ip string) bool {
	return regexpx.IsMatchString(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`, ip)
}

// ConvertIpToUint32 converts ip address to an uint32 integer.
func ConvertIpToUint32(ip string) uint32 {
	netIp := net.ParseIP(ip)
	if netIp == nil {
		return 0
	}
	return binary.BigEndian.Uint32(netIp.To4())
}

// ConvertUint32ToIp converts an uint32 integer ip address to its string type address.
func ConvertUint32ToIp(long uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, long)
	return net.IP(ipByte).String()
}

// ParseAddress parses `address` to its ip and port.
// Eg: 192.168.1.1:80 -> 192.168.1.1, 80
func ParseAddress(address string) (string, int) {
	match, err := regexpx.MatchString(`^(.+):(\d+)$`, address)
	if err == nil {
		i, _ := strconv.Atoi(match[2])
		return match[1], i
	}
	return "", 0
}

// GetSegment returns the segment of given ip address.
// Eg: 192.168.2.102 -> 192.168.2
func GetSegment(ip string) string {
	match, err := regexpx.MatchString(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`, ip)
	if err != nil || len(match) < 4 {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s", match[1], match[2], match[3])
}

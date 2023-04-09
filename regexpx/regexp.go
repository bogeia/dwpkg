package regexpx

// IsMatch checks whether given bytes `src` matches `pattern`.
func IsMatch(pattern string, src []byte) bool {
	if r, err := getRegexp(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

// IsMatchString checks whether given string `src` matches `pattern`.
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// Match return bytes slice that matched `pattern`.
func Match(pattern string, src []byte) ([][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchString return strings that matched `pattern`.
func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

// MatchAll return all bytes slices that matched `pattern`.
func MatchAll(pattern string, src []byte) ([][][]byte, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

// MatchAllString return all strings that matched `pattern`.
func MatchAllString(pattern string, src string) ([][]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindAllStringSubmatch(src, -1), nil
	} else {
		return nil, err
	}
}

package ptt

import "strconv"

type ver struct {
	v     string
	i     int
	major string
	minor string
	patch string
}

var v = ver{
	v:     "0.3.0", // x-release-please-version
	major: "0",     // x-release-please-major
	minor: "3",     // x-release-please-minor
	patch: "0",     // x-release-please-patch
}

func (v ver) Int() int {
	if v.i != 0 {
		return v.i
	}
	major, _ := strconv.Atoi(non_digits_regex.ReplaceAllLiteralString(v.major, ""))
	minor, _ := strconv.Atoi(non_digits_regex.ReplaceAllLiteralString(v.minor, ""))
	patch, _ := strconv.Atoi(non_digits_regex.ReplaceAllLiteralString(v.patch, ""))
	v.i = major*1000000 + minor*1000 + patch
	return v.i
}

func (v ver) String() string {
	return v.v
}

func Version() ver {
	return v
}

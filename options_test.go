package jettison

import (
	"testing"
)

func Test_hasKeypathPrefix(t *testing.T) {
	samples := []struct {
		ss [2][]string
		want bool
	} {
		{ss: [2][]string{{"key1","subkey"}, {"key1"}}, want: true},
		{ss: [2][]string{{"key1"}, {"key1","subkey"}}, want: false},
		{ss: [2][]string{{"key1", "subkey_1"}, {"key1","subkey"}}, want: false},
		{ss: [2][]string{{"key1","key2", "leaf"}, {"key1", "key2"}}, want: true},
	}
	for i, s := range samples {
		b := hasKeyPathPrefix(s.ss[0],s.ss[1])
		if s.want != b {
			t.Errorf("%d: got = %t, want = %t", i, b, s.want)
		}
	}
}

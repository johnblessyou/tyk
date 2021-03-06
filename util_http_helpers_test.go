package main

import (
	"net/http"
	"testing"
)

func TestGetIPFromRequest(t *testing.T) {
	tests := []struct {
		remote, forwarded, want string
	}{
		// missing ip or port
		{"", "", ""},
		{":80", "", ""},
		{"1.2.3.4", "", ""},
		{"[::1]", "", ""},
		// not forwarded
		{"1.2.3.4:80", "", "1.2.3.4"},
		{"[::1]:80", "", "::1"},
		// forwarded
		{"1.2.3.4:80", "5.6.7.8, px1, px2", "5.6.7.8"},
		{"[::1]:80", "::2", "::2"},
	}
	for _, tc := range tests {
		r := &http.Request{RemoteAddr: tc.remote, Header: http.Header{}}
		r.Header.Set("x-forwarded-for", tc.forwarded)
		got := GetIPFromRequest(r)
		if got != tc.want {
			t.Errorf("GetIPFromRequest({%q, %q}) got %q, want %q",
				tc.remote, tc.forwarded, got, tc.want)
		}
	}
}

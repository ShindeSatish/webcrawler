package crawler

import (
	"testing"
)

func TestResolveLink(t *testing.T) {
	tests := []struct {
		base, href, want string
	}{
		{"http://example.com", "/foo", "http://example.com/foo"},
		{"http://example.com/foo", "../bar", "http://example.com/bar"},
		{"http://example.com/foo/", "baz", "http://example.com/foo/baz"},
	}

	for _, tc := range tests {
		got := resolveLink(tc.href, tc.base)
		if got != tc.want {
			t.Errorf("resolveLink(%q, %q) = %q, want %q", tc.href, tc.base, got, tc.want)
		}
	}
}

func TestShouldCrawl(t *testing.T) {
	tests := []struct {
		link, rootURL string
		want          bool
	}{
		{"http://example.com/foo", "http://example.com", true},
		{"http://example.com/foo", "http://another.com", false},
		{"https://example.com/bar", "http://example.com", true},
	}

	for _, tc := range tests {
		got := shouldCrawl(tc.link, tc.rootURL)
		if got != tc.want {
			t.Errorf("shouldCrawl(%q, %q) = %v, want %v", tc.link, tc.rootURL, got, tc.want)
		}
	}
}

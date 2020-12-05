package types

import "net/url"

// URL is url.URL
type URL = url.URL

// ParseURL is url.Parse
func ParseURL(string string) (*URL, error) { return url.Parse(string) }

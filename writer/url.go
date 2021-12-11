package writer

import (
	"net/url"
	"strings"
)

const (
	labelKey = "__label"
)

// WithLabel adds a label to the URL.
func WithLabel(raw, label string) string {
	u, err := url.Parse(raw)
	if err != nil {
		panic(err)
	}

	q := u.Query()
	q.Set(labelKey, label)
	u.RawQuery = q.Encode()

	return u.String()
}

// ExtractLabel extracts the label from the URL.
func ExtractLabel(s string) (raw, label string, ok bool) {
	u, err := url.Parse(strings.TrimSpace(s))
	if err != nil {
		return
	}

	q := u.Query()
	label = q.Get(labelKey)
	if label == "" {
		return
	}

	q.Del(labelKey)
	u.RawQuery = q.Encode()
	raw = u.String()
	ok = true
	return
}

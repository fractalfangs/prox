package safemode

import (
	"html"
	"strings"

	"github.com/router-for-me/CLIProxyAPI/v8/internal/webstyle"
)

var exampleAPIKeys = map[string]struct{}{
	"your-api-key-1": {},
	"your-api-key-2": {},
	"your-api-key-3": {},
}

// ExampleAPIKeys returns configured top-level API keys that still use template values.
func ExampleAPIKeys(keys []string) []string {
	if len(keys) == 0 {
		return nil
	}

	matches := make([]string, 0, len(keys))
	seen := make(map[string]struct{}, len(exampleAPIKeys))
	for _, key := range keys {
		trimmed := strings.TrimSpace(key)
		if _, ok := exampleAPIKeys[trimmed]; !ok {
			continue
		}
		if _, exists := seen[trimmed]; exists {
			continue
		}
		seen[trimmed] = struct{}{}
		matches = append(matches, trimmed)
	}
	if len(matches) == 0 {
		return nil
	}
	return matches
}

// HasExampleAPIKeys reports whether any configured top-level API key is a template value.
func HasExampleAPIKeys(keys []string) bool {
	return len(ExampleAPIKeys(keys)) > 0
}

// ExampleAPIKeyWarningPageHTML returns the setup warning page HTML.
func ExampleAPIKeyWarningPageHTML(keys []string, managementPath string) string {
	var b strings.Builder
	b.WriteString(`<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Example API key detected</title><style>` + webstyle.CSS + `</style></head><body><main><span class="brand">cliproxyapi</span><p class="tags"><b>safe mode</b> ⋅ proxy endpoints disabled</p><h1>Example API key detected</h1><p>Proxy API endpoints are disabled because the top-level <code>api-keys</code> configuration still contains template values.</p>`)
	if len(keys) > 0 {
		b.WriteString(`<p>Replace these values before using the proxy:</p><ul>`)
		for _, key := range keys {
			b.WriteString(`<li><code>`)
			b.WriteString(html.EscapeString(key))
			b.WriteString(`</code></li>`)
		}
		b.WriteString(`</ul>`)
	}
	b.WriteString(`<p>Set strong random API keys, then retry the proxy endpoint.</p>`)
	if trimmed := strings.TrimSpace(managementPath); trimmed != "" {
		b.WriteString(`<nav class="nav"><a href="`)
		b.WriteString(html.EscapeString(trimmed))
		b.WriteString(`">Open Management</a></nav>`)
	}
	b.WriteString(`</main></body></html>`)
	return b.String()
}

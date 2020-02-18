// Package setcookie is a small package to parse Set-Cookie header
// value containing multiple cookies separated by comma.
//
// Based on similar code in Java and Javascript:
// - https://github.com/google/j2objc/commit/16820fdbc8f76ca0c33472810ce0cb03d20efe25
// - https://github.com/nfriedly/set-cookie-parser/blob/master/lib/set-cookie.js
//
// License: MIT
package setcookie

// SplitCookies splits Set-Cookie string containing multiple cookies separated by comma.
// Only a comma before KEY=VALUE of the next cookie can cause the split.
func SplitCookies(setCookieStr string) (outArr []string) {
	pos := 0
	inputLen := len(setCookieStr)

	// skips whitespace, returning true if there are more chars to read
	skipSpace := func() bool {
		for pos < inputLen && (setCookieStr[pos] == '\t' || setCookieStr[pos] == ' ') {
			pos++
		}
		return pos < inputLen
	}

	hasNext := func() bool {
		return pos < len(setCookieStr)
	}

	next := func() string {
		start := pos
		for skipSpace() {
			if setCookieStr[pos] == ',' {
				lastComma := pos
				pos++
				skipSpace()
				nextStart := pos
				for pos < inputLen && setCookieStr[pos] != '=' && setCookieStr[pos] != ';' && setCookieStr[pos] != ',' {
					pos++
				}
				if pos < inputLen && setCookieStr[pos] == '=' {
					// pos is inside the next cookie, so back up and return it
					pos = nextStart
					return setCookieStr[start:lastComma]
				}
				pos = lastComma
			}
			pos++
		}
		return setCookieStr[start:]
	}

	for hasNext() {
		outArr = append(outArr, next())
	}

	return
}

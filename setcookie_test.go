package setcookie

import (
	"strings"
	"testing"
)

func TestEmpty(t *testing.T) {
	setCookie := ""
	arr := SplitCookies(setCookie)

	if len(arr) != 0 {
		t.Fatal("empty array expected for an empty setCookie string")
	}
}

func TestSingle(t *testing.T) {
	setCookie := `SIDCC=XXX; path=/; domain=.google.com; priority=high`
	arr := SplitCookies(setCookie)

	if len(arr) != 1 {
		t.Fatal("expected single-element array for a single set-cookie string")
	}

	if arr[0] != setCookie {
		t.Fatal("parsed cookie without any whitespaces around it must not be modified")
	}
}

func TestSingleExpires(t *testing.T) {
	setCookie := `SIDCC=XXX; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/; domain=.google.com; priority=high`
	arr := SplitCookies(setCookie)

	if len(arr) != 1 {
		t.Fatal("expected single-element array for a single set-cookie string")
	}

	if arr[0] != setCookie {
		t.Fatal("parsed cookie without any whitespaces around it must not be modified")
	}
}

func TestMultiple(t *testing.T) {
	setCookie := `SIDCC=XXX; path=/; domain=.google.com; priority=high, BBxB=second; path=/`
	arr := SplitCookies(setCookie)

	if len(arr) != 2 {
		t.Fatal("expected two-element array for a set-cookie string with two comma-separated cookies")
	}

	if arr[0] != `SIDCC=XXX; path=/; domain=.google.com; priority=high` {
		t.Fatal("first cookie not parsed correctly")
	}

	if arr[1] != `BBxB=second; path=/` {
		t.Fatal("second cookie not parsed correctly")
	}
}

func TestMultipleWhitespace(t *testing.T) {
	setCookie := `SIDCC=XXX; path=/; domain=.google.com; priority=high  ,  BBxB=second; path=/`
	arr := SplitCookies(setCookie)

	if len(arr) != 2 {
		t.Fatal("expected two-element array for a set-cookie string with two comma-separated cookies")
	}

	if strings.TrimSpace(arr[0]) != `SIDCC=XXX; path=/; domain=.google.com; priority=high` {
		t.Fatal("first cookie not parsed correctly")
	}

	if strings.TrimSpace(arr[1]) != `BBxB=second; path=/` {
		t.Fatal("second cookie not parsed correctly")
	}
}

func TestMultipleExpires(t *testing.T) {
	setCookie := `SIDCC=XXX; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/; domain=.google.com; priority=high, BBB=second; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/`
	arr := SplitCookies(setCookie)

	if len(arr) != 2 {
		t.Fatal("expected two-element array for a set-cookie string with two comma-separated cookies")
	}

	if arr[0] != `SIDCC=XXX; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/; domain=.google.com; priority=high` {
		t.Fatal("first cookie not parsed correctly")
	}

	if arr[1] != `BBB=second; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/` {
		t.Fatal("second cookie not parsed correctly")
	}
}

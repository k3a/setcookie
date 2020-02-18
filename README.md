[![GoDoc](https://godoc.org/github.com/k3a/setcookie?status.svg)](https://godoc.org/github.com/k3a/setcookie)
[![Build Status](https://travis-ci.org/k3a/setcookie.svg?branch=master)](https://travis-ci.org/k3a/setcookie)
[![Coverage Status](https://coveralls.io/repos/k3a/setcookie/badge.svg?branch=master&service=github)](https://coveralls.io/github/k3a/setcookie?branch=master)
[![Report Card](https://goreportcard.com/badge/github.com/k3a/setcookie)](https://goreportcard.com/report/github.com/k3a/setcookie)

# Set-Cookie Parser

This code

```go
package main

import (
	"fmt"

	"github.com/k3a/setcookie"
)

func main() {
	setCookieStr := `AAA=XXX; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/; priority=high, BBB=YYY; expires=Wed, 17-Feb-2021 19:15:01 GMT; path=/`
	arr := setcookie.SplitCookies(setCookie)
	fmt.Printf("I hold %d cookies!\n", len(arr))
}
```

prints

```
I hold 2 cookies!
```

Because the string variable `setCookieStr` contains two comma-separated cookies. 8-)

The code is MIT-licensed and well-tested.

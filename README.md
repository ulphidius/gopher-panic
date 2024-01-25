# gopherpanic

[![push_master](https://github.com/ulphidius/gopherpanic/actions/workflows/push_master.yml/badge.svg)](https://github.com/ulphidius/gopherpanic/actions/workflows/push_master.yml)
[![codecov](https://codecov.io/gh/ulphidius/gopherpanic/branch/master/graph/badge.svg?token=mN8q98Z1kU)](https://codecov.io/gh/ulphidius/gopherpanic)
[![Go Report Card](https://goreportcard.com/badge/github.com/ulphidius/gopherpanic)](https://goreportcard.com/report/github.com/ulphidius/gopherpanic)
[![Go Reference](https://pkg.go.dev/badge/github.com/ulphidius/gopherpanic.svg)](https://pkg.go.dev/github.com/ulphidius/gopherpanic)

gopher-panic has for aims to enhance the go error system


## Install

```sh
go get github.com/ulphidius/gopherpanic
```

## Configuration

**GOPHERPANIC_FORMAT** allow you to change the *Error* function output.
It's a numeric value which represent the Format type.

- GNU Format: 0
- GNU Format with traces: 1
- Custom Format: 2
- Custom Format with traces: 3

## Example

```go
package main

import (
	"fmt"

	"github.com/ulphidius/gopherpanic"
)

func main() {
	fmt.Println(gopherpanic.GopherpanicFormat)
	result, err := div(10, 0)
	if err != nil {
		panic(gopherpanic.Wrap(gopherpanic.ClientError, "fail to compute", err))
	}

	fmt.Println(result)
}

func div(x, y int64) (int64, *gopherpanic.Error) {
	if y == 0 {
		return 0, gopherpanic.New(gopherpanic.ClientError, "cannot divide by 0")
	}

	return x / y, nil
}
```

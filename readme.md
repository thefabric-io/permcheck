# permcheck

[![Go Report Card](https://goreportcard.com/badge/github.com/thefabric-io/permcheck)](https://goreportcard.com/report/github.com/thefabric-io/permcheck)

`permcheck` is a Golang package that provides a convenient way to check and manage permissions. It allows permissions to be combined using boolean logic (AND, OR), making it a flexible tool for implementing authorization systems.

## Features

- **Simple to use**: Easily define and check permissions.
- **Logical operators**: Combine permissions with AND and OR operations.
- **Customizable**: Define your own permission types.

## Installation

To install permcheck, use `go get`:

```bash
go get github.com/thefabric-io/permcheck
```

## Usage

```go
package main

import (
	"errors"
	"fmt"
	
	"github.com/thefabric-io/permcheck"
)

func main() {
	// Create new permissions
	p1 := permcheck.New("read", errors.New("read access required"))
	p2 := permcheck.New("write", errors.New("write access required"))

	// Create a combined permission
	p3 := permcheck.And(p1, p2)

	// Check permissions
	err := p3.Satisfies([]string{"read", "write"})
	if err != nil {
		fmt.Println(err)
	}
}
```

## Contributing

Contributions to `permcheck` are welcome! Please see [CONTRIBUTING.md](https://github.com/thefabric-io/permcheck/contributing.md) for details on how to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/thefabric-io/permcheck/license.md) file for details.

---

We hope this library proves to be a valuable tool in your software development toolkit.
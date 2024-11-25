[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/tern/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/tern/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/tern)](https://pkg.go.dev/github.com/yyle88/tern)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/tern/master.svg)](https://coveralls.io/github/yyle88/tern?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/tern.svg)](https://github.com/yyle88/tern/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/tern)](https://goreportcard.com/report/github.com/yyle88/tern)

# tern

**`tern`** is a lightweight and versatile Go package designed to simplify conditional logic using concise ternary expressions. With `tern`, you can write clear and expressive code that is easy to read and maintain.

## README

[中文说明](README.zh.md)

## Features

- **Generic Support**: Fully utilizes Go’s generics, making it type-safe and reusable across a wide range of data types.
- **Flexible Logic**: Supports both direct boolean conditions and conditional functions.
- **Deferred Execution**: Allows lazy evaluation using functions to compute values only when needed, improving performance in certain cases.
- **Zero Value Handling**: Provides utilities to return zero values when no fallback is provided.

## Installation

```sh
go get github.com/yyle88/tern
```

## Usage

The `tern` package offers various helper functions for different conditional scenarios:

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	// Simple conditional expression
	result := tern.BVV(true, "A", "B")
	fmt.Println(result) // Output: A

	// Deferred execution for the second value
	result = tern.BVF(false, "A", func() string { return "Computed B" })
	fmt.Println(result) // Output: Computed B

	// Using zero values as fallback
	result = tern.BV(false, "A")
	fmt.Println(result) // Output: (empty string)
}
```

### Available Functions

The `tern` package provides the following functions based on condition types and value evaluation methods:

| **Function** | **Condition Type** | **First Value**          | **Second Value**         |
|--------------|--------------------|--------------------------|--------------------------|
| `BVV`        | `bool`             | Direct value             | Direct value             |
| `BVF`        | `bool`             | Direct value             | Function returning value |
| `BFV`        | `bool`             | Function returning value | Direct value             |
| `BFF`        | `bool`             | Function returning value | Function returning value |
| `FVV`        | `func() bool`      | Direct value             | Direct value             |
| `FVF`        | `func() bool`      | Direct value             | Function returning value |
| `FFV`        | `func() bool`      | Function returning value | Direct value             |
| `FFF`        | `func() bool`      | Function returning value | Function returning value |

Additionally, utility functions like `BV`, `BF`, `FV`, and `FF` simplify zero-value fallbacks.

### Lazy Evaluation Example

Using deferred execution ensures unnecessary computations are avoided:

```go
func computeHeavyValue() string {
	fmt.Println("Heavy computation...")
	return "Heavy Result"
}

result := tern.BVF(false, "Default", computeHeavyValue)
// Output: Default (computeHeavyValue() is not executed)
```

### Custom Zero Values

The package uses Go’s `Zero[T]()`, which automatically provides the zero value for any type:

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	fmt.Println(tern.Zero[int]())    // Output: 0
	fmt.Println(tern.Zero[string]()) // Output: (empty string)
}
```

## Why Use `tern`?

1. **Readability**: Simplifies conditional logic, making the code easier to follow.
2. **Flexibility**: Works seamlessly with both direct values and deferred computation.
3. **Performance**: Avoids unnecessary computation by using lazy evaluation.
4. **Generics**: Leverages Go’s generics for maximum flexibility and type safety.

## Contributing

Contributions are welcome! If you’d like to improve the package or fix a bug, feel free to open an issue or submit a pull request on [GitHub](https://github.com/yyle88/tern).

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

## Contributing

Feel free to contribute or improve the package! Stars and pull requests are always welcome!

Thank you for using `tern`!

---

## See stars

[![see stars](https://starchart.cc/yyle88/tern.svg?variant=adaptive)](https://starchart.cc/yyle88/tern)

Give me stars! Thank you!!!

---

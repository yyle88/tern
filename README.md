<p align="center">
  <img 
    alt="golang-tern-expression logo" 
    src="assets/golang-tern-logo.jpeg" 
    style="max-height: 500px; width: auto; max-width: 100%;" 
  />
</p>
<h3 align="center">golang-tern-expression</h3>
<p align="center">simple tern expression <code>condition ? expression 1 : expression 2</code> in golang</p>

---

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/tern/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/tern/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/tern)](https://pkg.go.dev/github.com/yyle88/tern)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/tern/master.svg)](https://coveralls.io/github/yyle88/tern?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/tern.svg)](https://github.com/yyle88/tern/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/tern)](https://goreportcard.com/report/github.com/yyle88/tern)

# tern

**`tern`** is a lightweight and versatile Go package designed to streamline conditional logic with concise ternary expressions, helping you write clean, expressive, and maintainable code with ease.

## README

[中文说明](README.zh.md)

## Features

- **Generic Support**: Fully leverages Go’s generics for type safety and flexibility across various data types.
- **Flexible Logic**: Provides robust support for both boolean conditions and conditional functions.
- **Lazy Evaluation**: Optimizes performance by computing values only when necessary through deferred execution.
- **Zero Value Handling**: Offers utilities to return default zero values for any type when no fallback is supplied.

## Installation

```sh
go get github.com/yyle88/tern
```

## Usage

The `tern` package provides multiple helper functions tailored to handle various conditional scenarios.

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	// Basic conditional selection
	result := tern.BVV(true, "Option A", "Option B")
	fmt.Println(result) // Output: Option A

	// Deferred execution for fallback value
	result = tern.BVF(false, "Default", func() string { return "Computed Fallback" })
	fmt.Println(result) // Output: Computed Fallback

	// Handling zero values as fallback
	result = tern.BV(false, "Fallback")
	fmt.Println(result) // Output: (empty string)
}
```

### Function Overview

Here is an overview of the functions provided by the `tern` package:

| **Function** | **Condition Type** | **Primary Value**        | **Fallback Value**       |
|--------------|--------------------|--------------------------|--------------------------|
| `BVV`        | `bool`             | Direct value             | Direct value             |
| `BVF`        | `bool`             | Direct value             | Function returning value |
| `BFV`        | `bool`             | Function returning value | Direct value             |
| `BFF`        | `bool`             | Function returning value | Function returning value |
| `FVV`        | `func() bool`      | Direct value             | Direct value             |
| `FVF`        | `func() bool`      | Direct value             | Function returning value |
| `FFV`        | `func() bool`      | Function returning value | Direct value             |
| `FFF`        | `func() bool`      | Function returning value | Function returning value |

### Lazy Evaluation Example

Deferred execution ensures unnecessary computations are avoided, making your code more efficient:

```go
func expensiveComputation() string {
	fmt.Println("Computing...")
	return "Heavy Result"
}

result := tern.BVF(false, "Default", expensiveComputation)
// Output: Default (expensiveComputation is not executed)
```

### Working with Zero Values

The package provides `Zero[T]()`, a utility that returns the zero value for any generic type:

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

---

### Handling Zero-Value Fallbacks

The package includes functions that automatically handle zero values when the condition is not met:

| **Function** | **Condition Type** | **Primary Value**        | **Fallback Value**       |
|--------------|--------------------|--------------------------|--------------------------|
| `BV`         | `bool`             | Direct value             | Zero value of type `T`   |
| `BF`         | `bool`             | Function returning value | Zero value of type `T`   |
| `FV`         | `func() bool`      | Direct value             | Zero value of type `T`   |
| `FF`         | `func() bool`      | Function returning value | Zero value of type `T`   |

---

### Additional Utilities in the `zerotern` Package

The `zerotern` subpackage extends `tern` with specialized utilities for comparing values to their zero value, adding more control for fallback scenarios.

| **Function** | **Comparison Type** | **Primary Value**        | **Fallback Value**       |
|--------------|---------------------|--------------------------|--------------------------|
| `VV`         | Direct comparison   | Direct value             | Direct value             |
| `VF`         | Direct comparison   | Direct value             | Function returning value |

#### Example: Using `VV` and `VF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	// Direct comparison
	result := zerotern.VV("non-zero", "fallback")
	fmt.Println(result) // Output: non-zero

	// Fallback with function
	result = zerotern.VF("", func() string { return "fallback func" })
	fmt.Println(result) // Output: fallback func
}
```

---

### Pointer-Based Utility Functions in `zerotern`

| **Function** | **Pointer Handling**    | **Fallback Value**       |
|--------------|-------------------------|--------------------------|
| `SetPV`      | Pointer to direct value | Direct value             |
| `SetPF`      | Pointer to direct value | Function returning value |

#### Example: Using `SetPV` and `SetPF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	var value int
	zerotern.SetPV(&value, 42)
	fmt.Println(value) // Output: 42

	value = 7
	zerotern.SetPF(&value, func() int { return 99 })
	fmt.Println(value) // Output: 7
}
```

---

### Why Choose `tern`?

1. **Improved Readability**: Simplifies conditional logic with concise and clear expressions.
2. **Performance Optimization**: Lazy evaluation avoids unnecessary computations.
3. **Type Safety**: Leverages Go’s generics for maximum flexibility and reliability.
4. **Versatility**: Supports a wide range of scenarios, including pointer handling and zero-value fallbacks.

## Contributing

Contributions are welcome! Feel free to report issues, suggest improvements, or submit pull requests on [GitHub](https://github.com/yyle88/tern).

## License

MIT License. See [LICENSE](LICENSE).

---

## Contributing

Feel free to contribute or improve the package! Stars and pull requests are always welcome!

Thank you for using `tern`!

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/tern.svg?variant=adaptive)](https://starchart.cc/yyle88/tern)

Give me stars! Thank you!!!

---

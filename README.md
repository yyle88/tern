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

以下是对 `BV`、`BF`、`FV` 和 `FF` 函数的详细表格说明：

---

### Utility Functions for Zero-Value Fallbacks

| **Function** | **Condition Type** | **Primary Value**        | **Fallback Value**     | **Description**                                                                                                                      |
|--------------|--------------------|--------------------------|------------------------|--------------------------------------------------------------------------------------------------------------------------------------|
| `BV`         | `bool`             | Direct value             | Zero value of type `T` | Returns the first value if the condition is true, otherwise returns the zero value of type `T`.                                      |
| `BF`         | `bool`             | Function returning value | Zero value of type `T` | Evaluates the function only if the condition is true, otherwise returns the zero value of type `T`.                                  |
| `FV`         | `func() bool`      | Direct value             | Zero value of type `T` | Evaluates the condition function and returns the first value if true, otherwise returns the zero value of type `T`.                  |
| `FF`         | `func() bool`      | Function returning value | Zero value of type `T` | Evaluates the condition function and returns the result of the first function if true, otherwise returns the zero value of type `T`. |

---

### Examples

#### Using `BV`

```go
result := tern.BV(false, "value")
// Output: (empty string, as the condition is false)
```

#### Using `BF`

```go
result := tern.BF(false, func() string { return "value" })
// Output: (empty string, as the condition is false)
```

#### Using `FV`

```go
condition := func() bool { return true }
result := tern.FV(condition, "value")
// Output: "value" (as the condition function returns true)
```

#### Using `FF`

```go
condition := func() bool { return false }
result := tern.FF(condition, func() string { return "value" })
// Output: (empty string, as the condition function returns false)
```

---

### Additional Utility Functions in `terngo` Package

The `terngo` subpackage extends the functionality of `tern` with additional utility functions, specifically designed for comparing values to their zero value. These functions allow for more expressive handling of default values when one of the inputs might be zero.

| **Function** | **Comparison Type** | **Primary Value** | **Fallback Value**       | **Description**                                                                                                                       |
|--------------|---------------------|-------------------|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| `VV`         | Direct comparison   | Direct value      | Direct value             | Returns the first value if it is not the zero value of type `T`, otherwise returns the second value.                                  |
| `VF`         | Direct comparison   | Direct value      | Function returning value | Returns the first value if it is not the zero value of type `T`, otherwise evaluates and returns the result of the fallback function. |

---

### Examples

#### Using `VV`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/terngo"
)

func main() {
	result := terngo.VV("non-zero", "fallback")
	fmt.Println(result) // Output: "non-zero"

	result = terngo.VV("", "fallback")
	fmt.Println(result) // Output: "fallback"
}
```

#### Using `VF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/terngo"
)

func main() {
	fallbackFunc := func() string { return "fallback from func" }

	result := terngo.VF("non-zero", fallbackFunc)
	fmt.Println(result) // Output: "non-zero"

	result = terngo.VF("", fallbackFunc)
	fmt.Println(result) // Output: "fallback from func"
}
```

---

### Key Benefits of `VV` and `VF`

- **Zero-Value Handling**: These functions leverage `tern.Zero[T]()` to check if the primary value is a zero value.
- **Fallback Options**: `VF` provides flexibility by allowing dynamic evaluation of the fallback value through a function.

---

By including these functions in the `terngo` package, developers gain even more expressive tools to simplify conditional logic in Go.

### Additional Utility Functions for Pointer Modification in `terngo`

The `terngo` subpackage introduces `PV` and `PF`, two specialized functions for safely modifying a pointer's value when it holds the zero value of its type. These functions ensure safer and more expressive manipulation of pointers while providing fallback options.

| **Function** | **Pointer Check**       | **Fallback Value**       | **Description**                                                                                                                       |
|--------------|-------------------------|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------|
| `PV`         | Pointer to direct value | Direct value             | Sets the pointed-to value to the fallback value if it is the zero value of type `T`. Panics if the pointer is `nil`.                  |
| `PF`         | Pointer to direct value | Function returning value | Sets the pointed-to value to the result of the fallback function if it is the zero value of type `T`. Panics if the pointer is `nil`. |

---

### Examples

#### Using `PV`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/terngo"
)

func main() {
	var value int
	terngo.PV(&value, 42)
	fmt.Println(value) // Output: 42

	value = 7
	terngo.PV(&value, 99)
	fmt.Println(value) // Output: 7
}
```

#### Using `PF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/terngo"
)

func main() {
	var value int
	fallbackFunc := func() int { return 42 }

	terngo.PF(&value, fallbackFunc)
	fmt.Println(value) // Output: 42

	value = 7
	terngo.PF(&value, fallbackFunc)
	fmt.Println(value) // Output: 7
}
```

---

### Key Features of `PV` and `PF`

- **Pointer Safety**: These functions panic if the provided pointer is `nil`, ensuring reliable execution.
- **Zero-Value Check**: Use `tern.Zero[T]()` to determine whether the current value is the zero value of its type.
- **Dynamic Fallback**: `PF` allows for fallback values to be generated dynamically via a function.

By leveraging `PV` and `PF`, developers can seamlessly handle zero-value assignments for pointer-based logic, maintaining both clarity and robustness in their Go code.

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

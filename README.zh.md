## tern

**`tern`** 是一个轻量级且多功能的 Go 包，通过简洁的三元表达式简化条件逻辑，让您可以轻松编写清晰、简洁、可维护的代码。

## 说明文档

[English README](README.md)

## 功能特点

- **泛型支持**：充分利用 Go 的泛型，类型安全且可重用，适用于各种数据类型。
- **灵活逻辑**：支持直接的布尔条件和基于函数的条件判断。
- **延迟执行**：支持使用函数延迟计算值，仅在需要时计算，有助于提高性能。
- **零值处理**：提供实用函数，在没有备用值时返回类型的零值。

## 安装

```sh
go get github.com/yyle88/tern
```

## 用法

`tern` 包为不同条件场景提供了多种辅助函数：

### 基本用法

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	// 简单的条件表达式
	result := tern.BVV(true, "A", "B")
	fmt.Println(result) // 输出: A

	// 使用延迟计算的第二个值
	result = tern.BVF(false, "A", func() string { return "计算后的B" })
	fmt.Println(result) // 输出: 计算后的B

	// 使用零值作为备用值
	result = tern.BV(false, "A")
	fmt.Println(result) // 输出: （空字符串）
}
```

### 可用函数

以下表格列出了 `tern` 提供的函数，按条件类型和值的计算方法分类：

| **函数** | **条件类型**      | **第一个值** | **第二个值** |
|--------|---------------|----------|----------|
| `BVV`  | `bool`        | 直接值      | 直接值      |
| `BVF`  | `bool`        | 直接值      | 函数返回值    |
| `BFV`  | `bool`        | 函数返回值    | 直接值      |
| `BFF`  | `bool`        | 函数返回值    | 函数返回值    |
| `FVV`  | `func() bool` | 直接值      | 直接值      |
| `FVF`  | `func() bool` | 直接值      | 函数返回值    |
| `FFV`  | `func() bool` | 函数返回值    | 直接值      |
| `FFF`  | `func() bool` | 函数返回值    | 函数返回值    |

### 延迟计算示例

使用延迟计算可以避免不必要的开销：

```go
func computeHeavyValue() string {
	fmt.Println("执行复杂计算...")
	return "计算结果"
}

result := tern.BVF(false, "默认值", computeHeavyValue)
// 输出: 默认值（computeHeavyValue() 未被执行）
```

### 自定义零值

`tern` 使用 Go 的 `Zero[T]()` 方法，根据类型自动提供零值：

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	fmt.Println(tern.Zero[int]())    // 输出: 0
	fmt.Println(tern.Zero[string]()) // 输出: （空字符串）
}
```

---

### 基于零值的备用处理函数

| **函数** | **条件类型**      | **主值** | **备用值** | **描述**                               |
|--------|---------------|--------|---------|--------------------------------------|
| `BV`   | `bool`        | 直接值    | 类型的零值   | 如果条件为真，则返回第一个值；否则返回类型的零值。            |
| `BF`   | `bool`        | 函数返回值  | 类型的零值   | 如果条件为真，则计算函数并返回结果；否则返回类型的零值。         |
| `FV`   | `func() bool` | 直接值    | 类型的零值   | 计算条件函数，如果为真，则返回第一个值；否则返回类型的零值。       |
| `FF`   | `func() bool` | 函数返回值  | 类型的零值   | 计算条件函数，如果为真，则计算第一个函数并返回结果；否则返回类型的零值。 |

---

### 示例

#### 使用 `BV`

```go
result := tern.BV(false, "值")
// 输出: （空字符串，因为条件为假）
```

#### 使用 `BF`

```go
result := tern.BF(false, func() string { return "值" })
// 输出: （空字符串，因为条件为假）
```

#### 使用 `FV`

```go
condition := func() bool { return true }
result := tern.FV(condition, "值")
// 输出: "值"（因为条件函数返回 true）
```

#### 使用 `FF`

```go
condition := func() bool { return false }
result := tern.FF(condition, func() string { return "值" })
// 输出: （空字符串，因为条件函数返回 false）
```

---

### `zerotern` 子包中的附加工具函数

`zerotern` 子包扩展了 `tern` 的功能，提供了更多与零值比较相关的实用函数。这些函数支持在输入可能为零值时更清晰地处理默认值。

| **函数** | **比较类型** | **主值** | **备用值** | **描述**                             |
|--------|----------|--------|---------|------------------------------------|
| `VV`   | 直接比较     | 直接值    | 直接值     | 如果第一个值不是类型的零值，则返回第一个值；否则返回第二个值。    |
| `VF`   | 直接比较     | 直接值    | 函数返回值   | 如果第一个值不是类型的零值，则返回第一个值；否则计算函数并返回结果。 |

---

### 示例

#### 使用 `VV`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	result := zerotern.VV("非零值", "备用值")
	fmt.Println(result) // 输出: "非零值"

	result = zerotern.VV("", "备用值")
	fmt.Println(result) // 输出: "备用值"
}
```

#### 使用 `VF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	fallbackFunc := func() string { return "函数计算的备用值" }

	result := zerotern.VF("非零值", fallbackFunc)
	fmt.Println(result) // 输出: "非零值"

	result = zerotern.VF("", fallbackFunc)
	fmt.Println(result) // 输出: "函数计算的备用值"
}
```

---

### 基于指针的工具函数 (`zerotern` 子包)

`zerotern` 子包提供了一组基于指针的赋值工具函数，用于检查指针内容是否为零值，并根据需要更新其内容。这些函数可以简化对指针类型值的默认处理逻辑。

| **函数** | **参数类型** | **主值** | **备用值** | **描述**                        |
|--------|----------|--------|---------|-------------------------------|
| `PV`   | 指针及备用值   | 指针指向的值 | 直接值     | 如果指针内容为零值，则将指针内容更新为备用值。       |
| `PF`   | 指针及备用值   | 指针指向的值 | 函数返回值   | 如果指针内容为零值，则通过函数计算备用值，再更新指针内容。 |

---

### 示例

#### 使用 `PV`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	value := ""
	ptr := &value

	zerotern.PV(ptr, "备用值")
	fmt.Println(*ptr) // 输出: "备用值"

	zerotern.PV(ptr, "新值")
	fmt.Println(*ptr) // 输出: "备用值"（因为内容已经非零值）
}
```

#### 使用 `PF`

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern/zerotern"
)

func main() {
	value := ""
	ptr := &value

	fallbackFunc := func() string { return "函数计算的备用值" }

	zerotern.PF(ptr, fallbackFunc)
	fmt.Println(*ptr) // 输出: "函数计算的备用值"

	zerotern.PF(ptr, func() string { return "新计算值" })
	fmt.Println(*ptr) // 输出: "函数计算的备用值"（因为内容已经非零值）
}
```

---

通过 `PV` 和 `PF`，`zerotern` 子包可以有效地处理指针内容的零值逻辑，提供了一种简洁而高效的方式来实现默认值赋值。

---

## `tern` 有什么优势？

1. **可读性**：简化条件逻辑，代码更加易读。
2. **灵活性**：同时支持直接值和延迟计算。
3. **性能优化**：通过延迟计算避免不必要的计算。
4. **泛型支持**：利用 Go 的泛型，实现灵活且类型安全的逻辑处理。

## 参与贡献

欢迎贡献代码！如果您有改进建议或发现了问题，可以在 [GitHub](https://github.com/yyle88/tern) 提交 Issue 或 Pull Request。

## 许可证

项目使用 MIT 许可证，详见 [LICENSE](LICENSE) 文件。 

## 其它

我的函数名很简单。

## 谢谢
有兴趣的可以试用。

希望大家给个星星。

# tern

**`tern`** 是一个轻量且灵活的 Go 包，用于通过简洁的三元表达式简化条件逻辑。使用 `tern`，您可以编写清晰且易于维护的代码。

## README

[ENGLISH-DOC](README.md)

## 功能特性

- **泛型支持**：充分利用 Go 的泛型功能，实现类型安全，可复用性强。
- **灵活的逻辑**：支持直接的布尔条件和基于函数的条件逻辑。
- **延迟执行**：支持使用函数进行懒加载，仅在需要时计算值，从而提升性能。
- **零值处理**：提供工具函数，当无后备值时返回类型的零值。

## 安装

```sh
go get github.com/yyle88/tern
```

## 使用说明

`tern` 包提供了针对不同条件场景的多种辅助函数：

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

	// 为第二个值使用延迟执行
	result = tern.BVF(false, "A", func() string { return "计算出的 B" })
	fmt.Println(result) // 输出: 计算出的 B

	// 使用零值作为后备值
	result = tern.BV(false, "A")
	fmt.Println(result) // 输出: (空字符串)
}
```

### 提供的函数

`tern` 包根据条件类型和值的计算方式提供以下函数：

| **函数名称** | **条件类型**      | **第一个值** | **第二个值** |
|----------|---------------|----------|----------|
| `BVV`    | `bool`        | 直接值      | 直接值      |
| `BVF`    | `bool`        | 直接值      | 返回值的函数   |
| `BFV`    | `bool`        | 返回值的函数   | 直接值      |
| `BFF`    | `bool`        | 返回值的函数   | 返回值的函数   |
| `FVV`    | `func() bool` | 直接值      | 直接值      |
| `FVF`    | `func() bool` | 直接值      | 返回值的函数   |
| `FFV`    | `func() bool` | 返回值的函数   | 直接值      |
| `FFF`    | `func() bool` | 返回值的函数   | 返回值的函数   |

此外，诸如 `BV`、`BF`、`FV` 和 `FF` 等工具函数用于简化返回零值的场景。

### 延迟计算示例

使用延迟执行可以避免不必要的计算：

```go
func computeHeavyValue() string {
	fmt.Println("执行重计算...")
	return "计算结果"
}

result := tern.BVF(false, "默认值", computeHeavyValue)
// 输出: 默认值 (computeHeavyValue() 不会被执行)
```

### 自定义零值

该包使用 Go 的 `Zero[T]()`，可自动为任何类型提供零值：

```go
package main

import (
	"fmt"
	"github.com/yyle88/tern"
)

func main() {
	fmt.Println(tern.Zero[int]())    // 输出: 0
	fmt.Println(tern.Zero[string]()) // 输出: (空字符串)
}
```

## `tern` 有什么优势？

1. **可读性**：简化条件逻辑，使代码更易理解。
2. **灵活性**：兼容直接值和延迟计算。
3. **性能**：通过懒加载避免不必要的计算。
4. **泛型支持**：利用 Go 泛型实现最大灵活性和类型安全。

## 贡献指南

欢迎贡献代码！如果您希望改进此包或修复 Bug，请随时在 [GitHub](https://github.com/yyle88/tern) 上提交 issue 或 pull request。

## 许可证

项目采用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。

## 其它

我的函数名很简单。

## 谢谢
有兴趣的可以试用。

希望大家给个星星。

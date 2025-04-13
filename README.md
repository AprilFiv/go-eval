# go - eval项目说明

## go - eval介绍
go - eval是一款极为纯粹的表达式引擎。通过简易的rete算法实现以及自底向上触发式的抽象语法树（AST）解释执行，具备以下显著特点：
- **交换律支持**：若表达式满足交换律，其执行顺序与输入顺序无关。例如，`A && B`和`B && A`完全等价。
- **部分错误容忍**：当表达式具有短路特性时，允许忽略错误以获取正确结果。如在`A && B`中，若`A`为`false`，即便`B`存在错误，仍能得出结果为`false`。
- **避免重复评估**：已运算过的表达式（前缀相关）无需再次评估。例如，`A + B + C`与`A + B`，公用前缀部分无需重复执行。

## 适用场景
鉴于上述特性，go - eval主要用于解决取数耗时问题，以减少表达式中变量的等待时间。典型应用场景如下：
- **风控规则引擎**：用于计算风控规则。
- **工作流系统**：用于确定条件分支。
- **营销发券规则**：在营销场景中决定发券条件。
- **其他类似场景**：……

## 使用说明
### 支持的数据类型
- `int64`
- `float64`
- `bool`
- `string`

### 支持的二元运算
- `+`、`-`、`*`、`/`、`%`
- `&&`、`||`
- `>=`、`<=`、`>`、`<`、`==`、`!=`

### 支持的一元运算
- `-`
- `!`

### 支持自定义函数


## 示例代码
```go
// 1.在同一映射中创建表达式
expressions := map[string]string{
    "Test1": "!A",
}

// 2. 解析表达式
descriptor, e := Parse(context.Background(), expressions)
if e != nil {
    panic(e)
}

// 3. 创建上下文
c := descriptor.Context()

// 4. 当参数可用时，通过“run”方法触发计算
descriptor.Run(context.Background(), c, "A", true, nil)

// 5. 查看输出，判断是否计算完成
if c.Output["Test1"] != nil {
    // 执行相关操作
}

// 额外示例：使用函数解析器进行解析
resolver := WithResolver(func(funcName string) (func(ctx context.Context, args ...interface{}) (interface{}, error), error) {
    if funcName != "myfunc" {
        return nil, errors.New("func error")
    }
    return func(ctx context.Context, args ...interface{}) (interface{}, error) {
        return args[0], nil
    }, nil
})
descriptor, e := Parse(context.Background(), expressions, resolver)
```

### 自定义函数支持
```go
expressions := map[string]string{
    "Test1": "myfunc(A)",
}
resolver := WithResolver(func(funcName string) (func(ctx context.Context, args ...interface{}) (interface{}, error), error) {
    if funcName != "myfunc" {
        return nil, errors.New("func error")
    }
    return func(ctx context.Context, args ...interface{}) (interface{}, error) {
        return args[0], nil
    }, nil
})
descriptor, e := Parse(context.Background(), expressions, resolver)
```


## Support
- **个人邮箱**：lyx51bb@gmail.com

本人在互联网风控开发领域拥有近10年工作经验。除本表达式引擎外，在风控领域的模型平台、指标系统、特征系统以及规则引擎等方面也有深入研究。欢迎联系我探讨相关问题。



你可以直接复制上述内容使用。要是后续还需要对这份README进行补充信息、调整格式等操作，随时都能找我 。 
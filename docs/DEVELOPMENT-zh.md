# 开发指南

这份指南是给想帮助改进 ccconfig 的贡献者的。

---

## 快速开始

### 前置要求

- 已安装 Go 1.21+
- 已安装 Git
- 基础 Go 知识

### 设置

```bash
# 克隆仓库
git clone https://github.com/jiangtao/cc-config.git
cd ccconfig

# 下载依赖
go mod download

# 构建
make build

# 安装到系统
make install
```

---

## 项目结构

```
ccconfig/
├── cmd/              # CLI 命令
│   ├── backup.go     # backup 命令
│   ├── restore.go    # restore 命令
│   ├── cache.go      # cache 命令
│   └── init.go       # init 命令
├── pkg/              # 核心包
│   ├── backup/       # 备份逻辑
│   ├── restore/      # 恢复逻辑
│   ├── config/       # 配置
│   ├── git/          # Git 操作
│   ├── i18n/         # 翻译
│   └── ui/           # 颜色和输出
├── web/              # 着陆页
└── install.sh        # 安装脚本
```

---

## 开发流程

### 修改代码

```bash
# 编辑代码
vim cmd/backup.go

# 本地运行
go run . backup --repo ~/test-config

# 格式化代码
make fmt

# 运行检查
make lint

# 运行测试
make test
```

### 常用命令

| 命令 | 说明 |
|------|------|
| `make build` | 构建二进制文件 |
| `make install` | 安装到系统 |
| `make test` | 运行所有测试 |
| `make fmt` | 格式化代码 |
| `make lint` | 运行代码检查 |
| `make tidy` | 清理依赖 |

---

## 添加功能

### 添加新命令

```go
// cmd/mycommand.go
package cmd

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "命令描述",
    RunE:  runMyCommand,
}

func init() {
    GetRootCommand().AddCommand(myCmd)
}

func runMyCommand(cmd *cobra.Command, args []string) error {
    // 你的代码在这里
    return nil
}
```

### 添加新语言

1. 创建 `pkg/i18n/{lang}.yaml`
2. 添加翻译

```yaml
# pkg/i18n/fr.yaml
my_message:
  title: "我的标题"
  description: "我的描述"
```

3. 在代码中使用：

```go
i18n.T("my_message.title", nil)
```

---

## 测试

### 编写测试

```go
func TestMyFunction(t *testing.T) {
    input := "test"
    expected := "期望输出"

    result := MyFunction(input)

    if result != expected {
        t.Errorf("期望 %v，得到 %v", expected, result)
    }
}
```

### 运行测试

```bash
# 运行所有测试
make test

# 运行特定包
go test ./pkg/backup -v

# 运行并查看覆盖率
go test -cover ./...
```

---

## 发布

### 创建发布

```bash
# 创建标签
git tag -a v1.0.0 -m "Release v1.0.0"

# 推送标签
git push origin v1.0.0
```

GitHub Actions 会自动：
- 运行测试
- 构建所有平台的二进制文件
- 创建 GitHub Release

---

## 贡献

我们欢迎贡献！

1. Fork 仓库
2. 创建分支
   ```bash
   git checkout -b feature/my-feature
   ```
3. 提交更改
   ```bash
   git commit -m "添加新功能"
   ```
4. 推送到你的 fork
   ```bash
   git push origin feature/my-feature
   ```
5. 创建 Pull Request

---

## 小贴士

### 代码风格

- 使用 `gofmt` 格式化代码
- 遵循 Go 约定
- 为导出函数添加注释

### 调试

```bash
# 详细输出运行
go run . backup --repo ~/test-config --verbose

# 使用 Delve 调试器
dlv debug ./... -- backup
```

---

## 需要帮助？

- 创建 issue: [https://github.com/jiangtao/cc-config/issues](https://github.com/jiangtao/cc-config/issues)
- 检查现有 issues: [https://github.com/jiangtao/cc-config/issues](https://github.com/jiangtao/cc-config/issues)

---

## 其他语言

- **English:** [DEVELOPMENT.md](DEVELOPMENT.md)

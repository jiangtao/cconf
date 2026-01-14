# GitHub Actions 自动恢复

本目录包含 GitHub Actions 工作流，用于自动构建和恢复 Claude Code 配置。

## 工作流说明

### 1. Build (`.github/workflows/build.yml`)

自动构建和测试 ccconfig，支持多平台编译。

**触发条件：**
- Push 到 main 分支
- 创建 tag (v*)
- Pull Request

**功能：**
- 运行测试
- 代码检查
- 跨平台构建 (darwin/amd64, darwin/arm64, linux/amd64, linux/arm64, windows/amd64)
- 自动创建 GitHub Release (tag 时)

### 2. Auto Restore (`.github/workflows/auto-restore.yml`)

自动恢复 Claude Code 配置。

**触发条件：**
- 手动触发 (workflow_dispatch)
- Push 到 main 分支

**配置：**

在你的仓库中设置以下 Secret：

| Secret | 说明 | 示例 |
|--------|------|------|
| `ANTHROPIC_AUTH_TOKEN` | Anthropic API Token | `sk-ant-...` |

**使用方式：**

1. **手动触发：**
   - 进入 Actions 页面
   - 选择 "Auto Restore Claude Config"
   - 点击 "Run workflow"
   - 可选：输入你的 ccconfig 仓库 URL

2. **自动触发：**
   - 每次代码合并到 main 分支时自动运行

### 3. PR Config Check (`.github/workflows/pr-comment.yml`)

在 PR 时检查配置状态并添加评论。

**触发条件：**
- PR 打开
- PR 更新

**配置：**

在你的仓库中设置以下 Secret：

| Secret | 说明 |
|--------|------|
| `CCCONFIG_REPO_URL` | 你的 ccconfig 配置仓库 URL |

## 快速开始

### 1. 克隆 ccconfig 仓库

```bash
git clone https://github.com/YOUR_USERNAME/ccconfig.git
cd ccconfig
```

### 2. 配置你的仓库 Secrets

在你的目标仓库中设置：

1. 进入 Settings → Secrets and variables → Actions
2. 添加以下 Secrets：
   - `ANTHROPIC_AUTH_TOKEN`: 你的 API Token
   - `CCCONFIG_REPO_URL`: `https://github.com/YOUR_USERNAME/ccconfig.git`

### 3. 复制工作流文件

将本目录中的 `.yml` 文件复制到你的仓库 `.github/workflows/` 目录。

### 4. 备份你的配置

```bash
cd ccconfig
./ccconfig backup
git add .
git commit -m "backup: Claude Code configuration"
git push
```

### 5. 触发自动恢复

**方式 1：自动触发**
```bash
# 合并到 main 分支后自动触发
git checkout main
git merge feature-branch
git push
```

**方式 2：手动触发**
1. 进入 GitHub Actions 页面
2. 选择 "Auto Restore Claude Config"
3. 点击 "Run workflow"

## 工作流输出

自动恢复完成后，会生成以下 Artifacts：

- `claude-config`: 完整的 Claude Code 配置目录

## 故障排除

### API Token 未设置

如果看到 `ANTHROPIC_AUTH_TOKEN not set` 错误：

1. 进入 Settings → Secrets and variables → Actions
2. 添加 `ANTHROPIC_AUTH_TOKEN` secret

### 配置仓库未找到

如果看到 `Config repository not found` 错误：

1. 检查 `CCCONFIG_REPO_URL` 是否正确
2. 确保仓库是公开的或已设置访问权限

### 自动恢复未生效

检查工作流日志：

1. 进入 Actions 页面
2. 选择失败的工作流运行
3. 查看详细日志

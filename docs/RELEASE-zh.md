# 发布指南

这份指南说明如何创建新的 ccconfig 发布版本。

---

## 版本号

我们使用语义化版本：`MAJOR.MINOR.PATCH`

- **MAJOR**: 重大变更
- **MINOR**: 新功能
- **PATCH**: Bug 修复

示例：
- `v1.0.0` - 首个稳定版本
- `v1.1.0` - 添加新功能
- `v1.1.1` - Bug 修复
- `v2.0.0` - 重大变更

---

## 如何发布

### 第一步：准备

```bash
# 确保所有测试通过
make test

# 如需要，更新 CHANGELOG.md
```

### 第二步：创建标签

```bash
# 创建并推送标签
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### 第三步：等待 GitHub Actions

推送标签后，GitHub Actions 会自动：

1. ✅ 运行测试
2. 🔨 为所有平台构建二进制文件
3. 📦 创建 GitHub Release
4. 📤 上传构建产物

这大约需要 5-10 分钟。

---

## 发布检查清单

发布前：
- [ ] 所有测试通过
- [ ] CHANGELOG.md 已更新
- [ ] 文档已更新
- [ ] 在 macOS 和 Linux 上测试过

发布后：
- [ ] 验证 GitHub Release 已创建
- [ ] 下载并测试二进制文件
- [ ] 如需要，更新 install.sh

---

## 紧急修复

如果在发布后发现严重 bug：

```bash
# 修复 bug
git commit -am "Hotfix: 严重 bug 修复"

# 创建新的补丁版本
git tag -a v1.0.1 -m "Hotfix: 严重 bug 修复"
git push origin main v1.0.1
```

---

## 需要帮助？

- 检查 [GitHub Actions](https://github.com/jiangtao/cc-config/actions)
- 创建 [issue](https://github.com/jiangtao/cc-config/issues)

---

## 其他语言

- **English:** [RELEASE.md](RELEASE.md)

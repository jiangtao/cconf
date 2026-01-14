# 发布指南

## 版本号规则

遵循语义化版本 (Semantic Versioning): `MAJOR.MINOR.PATCH`

- **MAJOR**: 不兼容的 API 变更
- **MINOR**: 向后兼容的功能新增
- **PATCH**: 向后兼容的问题修复

示例：
- `v1.0.0` - 首个稳定版本
- `v1.1.0` - 新增功能
- `v1.1.1` - Bug 修复
- `v2.0.0` - 重大变更

## 发布步骤

### 1. 更新版本

```bash
# 更新 go.mod 中的版本（如果有）
# 更新文档中的版本引用

# 创建并推送 tag
git tag -a v1.0.0 -m "Release v1.0.0

- Initial stable release
- Support backup/restore commands
- i18n support (en/zh)
- Cross-platform builds"
git push origin v1.0.0
```

### 2. GitHub Actions 自动构建

推送 tag 后，GitHub Actions 会：
1. 运行测试
2. 构建多平台二进制
3. 创建 GitHub Release
4. 上传构建产物

### 3. 验证 Release

1. 访问 GitHub Releases 页面
2. 下载各平台二进制
3. 测试基本功能

### 4. 发布公告

创建 Release Notes，包含：
- 变更内容
- 升级指南（如有）
- 已知问题

## 发布检查清单

- [ ] 所有测试通过
- [ ] 更新 CHANGELOG
- [ ] 更新文档中的版本引用
- [ ] 跨平台构建测试
- [ ] 创建 tag
- [ ] 推送 tag
- [ ] 验证 GitHub Release
- [ ] 发布公告

## 回滚计划

如果发布后发现严重问题：

```bash
# 删除远程 tag
git push origin :refs/tags/v1.0.0

# 删除本地 tag
git tag -d v1.0.0

# 创建修复版本
git tag -a v1.0.1 -m "Hotfix: ..."
git push origin v1.0.1
```

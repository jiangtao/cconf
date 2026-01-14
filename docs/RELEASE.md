# Release Guide

This guide explains how to create a new release of ccconfig.

---

## Version Numbers

We use Semantic Versioning: `MAJOR.MINOR.PATCH`

- **MAJOR**: Breaking changes
- **MINOR**: New features
- **PATCH**: Bug fixes

Examples:
- `v1.0.0` - First stable release
- `v1.1.0` - New feature added
- `v1.1.1` - Bug fix
- `v2.0.0` - Major changes

---

## How to Release

### Step 1: Prepare

```bash
# Make sure all tests pass
make test

# Update CHANGELOG.md if needed
```

### Step 2: Create Tag

```bash
# Create and push tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

### Step 3: Wait for GitHub Actions

After pushing the tag, GitHub Actions will automatically:

1. ✅ Run tests
2. 🔨 Build binaries for all platforms
3. 📦 Create GitHub Release
4. 📤 Upload build artifacts

This takes about 5-10 minutes.

---

## Release Checklist

Before releasing:
- [ ] All tests pass
- [ ] CHANGELOG.md updated
- [ ] Documentation updated
- [ ] Tested on macOS and Linux

After releasing:
- [ ] Verify GitHub Release created
- [ ] Download and test binaries
- [ ] Update install.sh if needed

---

## Hotfix

If you find a critical bug after release:

```bash
# Fix the bug
git commit -am "Hotfix: critical bug fix"

# Create new patch version
git tag -a v1.0.1 -m "Hotfix: critical bug fix"
git push origin main v1.0.1
```

---

## Need Help?

- Check [GitHub Actions](https://github.com/jiangtao/cc-config/actions)
- Create an [issue](https://github.com/jiangtao/cc-config/issues)

---

## Other Languages

- **中文文档:** [RELEASE-zh.md](RELEASE-zh.md)

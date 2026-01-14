# Quick Start: Deploy to Vercel

## Prerequisites
- Vercel account (sign up at https://vercel.com)
- GitHub repository access (jiangtao/cc-config)

## One-Line Deployment (After Vercel Login)

```bash
cd web && ./deploy.sh
```

## Manual Deployment Steps

### 1. Install Vercel CLI (if not already installed)
```bash
npm install -g vercel --registry https://registry.npmjs.org/
```

### 2. Login to Vercel
```bash
vercel login
```
This will open a browser for authentication.

### 3. Deploy
```bash
cd web
vercel --prod
```

Follow the prompts:
- Set project name: `ccconfig`
- Link to existing project? No (first deployment)
- Build settings: Use defaults from vercel.json

### 4. Note the deployment URL
Expected: `https://ccconfig.vercel.app` or similar

### 5. Verify deployment
```bash
curl https://ccconfig.vercel.app
```

Expected: HTML response with landing page content

## Alternative: Web Dashboard Deployment

1. Go to https://vercel.com/new
2. Click "Import Git Repository"
3. Select `jiangtao/cc-config`
4. Configure:
   - **Root Directory**: `web`
   - **Framework Preset**: Next.js
5. Click "Deploy"

## After Deployment

Once deployed, update the main README.md:

```markdown
[![Live Site](https://img.shields.io/badge/🔗-Live_Site-blue)](https://ccconfig.vercel.app)
```

Add this badge after the existing badges (line 6).

## Troubleshooting

**"No credentials found"**
- Run `vercel login` first
- Complete browser authentication

**Build fails**
- Verify local build: `npm run build`
- Check Vercel deployment logs

**Wrong version deployed**
- Vercel deploys latest git commit
- Ensure you've pushed to GitHub: `git push origin main`

## Continuous Deployment

After initial deployment, Vercel automatically deploys when you push to GitHub:

```bash
git add .
git commit -m "Update landing page"
git push origin main
```

## Current Status

✅ Ready for deployment
✅ Build tested locally
✅ Code pushed to GitHub
⏳ Awaiting Vercel authentication

For detailed documentation, see [DEPLOYMENT.md](./DEPLOYMENT.md).

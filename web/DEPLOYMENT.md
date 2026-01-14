# Vercel Deployment Guide

This guide explains how to deploy the ccconfig landing page to Vercel.

## Prerequisites

- Vercel account (sign up at https://vercel.com)
- GitHub repository with the landing page code
- Vercel CLI installed (optional, for command-line deployment)

## Deployment Options

### Option 1: Vercel Web Dashboard (Recommended)

1. **Push code to GitHub**
   ```bash
   git push origin main
   ```

2. **Import project in Vercel**
   - Go to https://vercel.com/new
   - Click "Import Git Repository"
   - Select your `cc-config` repository
   - Configure project settings:
     - **Framework Preset**: Next.js
     - **Root Directory**: `web`
     - **Build Command**: `npm run build`
     - **Output Directory**: `.next`
     - **Install Command**: `npm install`

3. **Deploy**
   - Click "Deploy"
   - Wait for deployment to complete (~2-3 minutes)
   - Vercel will provide a URL like `https://ccconfig.vercel.app`

4. **Configure custom domain** (optional)
   - Go to project Settings > Domains
   - Add your custom domain
   - Update DNS records as instructed

### Option 2: Vercel CLI

1. **Install Vercel CLI**
   ```bash
   npm install -g vercel
   ```

2. **Login to Vercel**
   ```bash
   vercel login
   ```
   This will open a browser for authentication.

3. **Deploy from web directory**
   ```bash
   cd web
   vercel --prod
   ```

4. **Follow the prompts**
   - Set project name: `ccconfig`
   - Link to existing project? No (first deployment)
   - Build settings: Use defaults from vercel.json

5. **Note the deployment URL**
   Expected: `https://ccconfig.vercel.app` or similar

## Deployment Configuration

The `web/vercel.json` file contains deployment settings:

```json
{
  "$schema": "https://openapi.vercel.sh/vercel.json",
  "framework": "nextjs",
  "buildCommand": "npm run build",
  "outputDirectory": ".next",
  "installCommand": "npm install",
  "devCommand": "npm run dev",
  "regions": ["sfo1"]
}
```

## Verification

After deployment, verify the site is live:

```bash
curl https://ccconfig.vercel.app
```

Expected: HTML response with landing page content

## Continuous Deployment

Vercel automatically deploys when you push to GitHub:

```bash
git add .
git commit -m "Update landing page"
git push origin main
```

Vercel will detect the push and automatically deploy the changes.

## Environment Variables

No environment variables are required for the landing page.

## Troubleshooting

**Build fails**
- Check that all dependencies are installed: `cd web && npm install`
- Verify build works locally: `npm run build`
- Check Vercel deployment logs for specific errors

**Deployment succeeds but site doesn't load**
- Check Vercel dashboard for deployment status
- Verify DNS propagation if using custom domain
- Check browser console for JavaScript errors

**Wrong version deployed**
- Vercel deploys the latest commit on the branch
- Check which branch is configured in Vercel project settings
- Verify you're pushing to the correct branch

## Performance Optimization

The landing page is optimized for performance:

- Static page generation (SSG)
- Optimized images
- CSS-in-JS with Tailwind CSS
- Minimal JavaScript bundle
- CDN delivery via Vercel Edge Network

Expected metrics:
- Lighthouse Performance Score: 95+
- First Contentful Paint: <1s
- Time to Interactive: <2s

## Monitoring

Monitor your deployment at:
- Vercel Dashboard: https://vercel.com/dashboard
- Project Analytics: https://vercel.com/[username]/ccconfig/analytics
- Deployment Logs: Available in each deployment details

## Rollback

If you need to rollback to a previous deployment:

1. Go to Vercel Dashboard > ccconfig > Deployments
2. Find the previous successful deployment
3. Click "Promote to Production"
4. Confirm the rollback

## Next Steps

After successful deployment:

1. Update README.md with the live site URL
2. Add the live site badge to documentation
3. Set up custom domain (optional)
4. Configure analytics (optional)
5. Set up automated testing (optional)

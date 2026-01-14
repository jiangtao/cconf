#!/bin/bash

# Vercel Deployment Script for ccconfig Landing Page
# This script deploys the landing page to Vercel

set -e

echo "🚀 ccconfig Landing Page Deployment Script"
echo "=========================================="

# Check if we're in the web directory
if [ ! -f "package.json" ]; then
    echo "❌ Error: Must be run from the web directory"
    echo "   Usage: cd web && ./deploy.sh"
    exit 1
fi

# Check if Vercel CLI is installed
if ! command -v vercel &> /dev/null; then
    echo "📦 Installing Vercel CLI..."
    npm install -g vercel --registry https://registry.npmjs.org/
fi

# Check if logged in to Vercel
if ! vercel whoami &> /dev/null; then
    echo "🔐 Please login to Vercel..."
    vercel login
fi

# Build the project locally first
echo "🔨 Building project locally..."
npm run build

# Deploy to Vercel
echo "🌐 Deploying to Vercel..."
vercel --prod

echo ""
echo "✅ Deployment complete!"
echo ""
echo "Next steps:"
echo "1. Visit the provided URL to verify deployment"
echo "2. Update ../README.md with the live site URL"
echo "3. Test the landing page functionality"

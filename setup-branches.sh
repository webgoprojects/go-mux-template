#!/bin/bash
# Script to create separate branches for each upgrade

set -e

# Initialize git if needed
if [ ! -d .git ]; then
    git init
    git config user.name "Developer"
    git config user.email "dev@example.com"
fi

# Commit initial state
git add -A
git commit -m "Initial commit: Basic Go Mux template" || echo "Already committed"

# Upgrade 1: Code Structure & Module Setup
git checkout -b upgrade-1-code-structure
# Apply upgrade 1 changes here (will be done manually)
git add -A
git commit -m "Upgrade 1: Fix code structure - create go.mod, wire handlers properly"
git push -u origin upgrade-1-code-structure || echo "Remote not configured, skipping push"

# Upgrade 2: Structured Logging
git checkout main
git checkout -b upgrade-2-structured-logging
# Apply upgrade 2 changes here
git add -A
git commit -m "Upgrade 2: Add structured logging with zap logger"
git push -u origin upgrade-2-structured-logging || echo "Remote not configured, skipping push"

# Upgrade 3: Environment Configuration
git checkout main
git checkout -b upgrade-3-config-management
# Apply upgrade 3 changes here
git add -A
git commit -m "Upgrade 3: Add environment configuration management"
git push -u origin upgrade-3-config-management || echo "Remote not configured, skipping push"

# Upgrade 4: Middleware
git checkout main
git checkout -b upgrade-4-middleware
# Apply upgrade 4 changes here
git add -A
git commit -m "Upgrade 4: Add middleware (CORS, request ID, logging)"
git push -u origin upgrade-4-middleware || echo "Remote not configured, skipping push"

# Upgrade 5: Health Check & Graceful Shutdown
git checkout main
git checkout -b upgrade-5-health-graceful-shutdown
# Apply upgrade 5 changes here
git add -A
git commit -m "Upgrade 5: Add health check endpoint and graceful shutdown"
git push -u origin upgrade-5-health-graceful-shutdown || echo "Remote not configured, skipping push"

# Final: README Update
git checkout main
git checkout -b upgrade-6-readme
# Apply README changes here
git add -A
git commit -m "Update README with all upgrades and documentation"
git push -u origin upgrade-6-readme || echo "Remote not configured, skipping push"

echo "All branches created successfully!"
echo "Note: You'll need to apply the actual code changes to each branch before committing"
